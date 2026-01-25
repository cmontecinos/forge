package cli

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

// PostCreateResult holds the results of post-create operations
type PostCreateResult struct {
	FrontendInstalled bool
	BackendInstalled  bool
	Duration          time.Duration
}

// RunPostCreate runs post-create hooks to install dependencies
func RunPostCreate(targetDir, stackType string) PostCreateResult {
	startTime := time.Now()
	result := PostCreateResult{}

	Info("Installing dependencies...")

	// Install frontend/mobile dependencies
	frontendDir := getFrontendDir(targetDir, stackType)
	if frontendDir != "" {
		if err := installNodeDeps(frontendDir); err != nil {
			Warn(fmt.Sprintf("Could not install frontend dependencies: %v", err))
		} else {
			result.FrontendInstalled = true
		}
	}

	// Install backend dependencies
	backendDir := filepath.Join(targetDir, "backend")
	if dirExists(backendDir) {
		if err := installGoDeps(backendDir); err != nil {
			Warn(fmt.Sprintf("Could not install backend dependencies: %v", err))
		} else {
			result.BackendInstalled = true
		}
	}

	result.Duration = time.Since(startTime)
	return result
}

// getFrontendDir returns the frontend directory based on stack type
func getFrontendDir(targetDir, stackType string) string {
	var dir string
	switch stackType {
	case "web":
		dir = filepath.Join(targetDir, "frontend")
	case "mobile":
		dir = filepath.Join(targetDir, "mobile")
	default:
		return ""
	}

	if dirExists(dir) {
		return dir
	}
	return ""
}

// installNodeDeps runs npm install in the given directory
func installNodeDeps(dir string) error {
	// Check if package.json exists
	packageJSON := filepath.Join(dir, "package.json")
	if _, err := os.Stat(packageJSON); os.IsNotExist(err) {
		return nil // No package.json, skip
	}

	// Prefer npm, fall back to yarn
	var cmd *exec.Cmd
	if commandExists("npm") {
		cmd = exec.Command("npm", "install", "--silent")
	} else if commandExists("yarn") {
		cmd = exec.Command("yarn", "install", "--silent")
	} else {
		return fmt.Errorf("neither npm nor yarn found")
	}

	cmd.Dir = dir
	cmd.Stdout = nil // Suppress output
	cmd.Stderr = nil

	return cmd.Run()
}

// installGoDeps runs go mod tidy in the given directory
func installGoDeps(dir string) error {
	// Check if go.mod exists
	goMod := filepath.Join(dir, "go.mod")
	if _, err := os.Stat(goMod); os.IsNotExist(err) {
		return nil // No go.mod, skip
	}

	if !commandExists("go") {
		return fmt.Errorf("go not found")
	}

	cmd := exec.Command("go", "mod", "tidy")
	cmd.Dir = dir
	cmd.Stdout = nil
	cmd.Stderr = nil

	return cmd.Run()
}

// commandExists checks if a command is available in PATH
func commandExists(name string) bool {
	_, err := exec.LookPath(name)
	return err == nil
}

// dirExists checks if a directory exists
func dirExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}
