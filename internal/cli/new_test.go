package cli

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestValidateProjectName tests project name validation
func TestValidateProjectName(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{"valid simple", "myapp", false},
		{"valid with hyphen", "my-app", false},
		{"valid with underscore", "my_app", false},
		{"valid with numbers", "myapp123", false},
		{"valid mixed", "My-App_123", false},
		{"empty", "", true},
		{"starts with number", "123app", true},
		{"starts with hyphen", "-myapp", true},
		{"contains space", "my app", true},
		{"contains special char", "my@app", true},
		{"too long", strings.Repeat("a", 65), true},
		{"max length", strings.Repeat("a", 64), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateProjectName(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("validateProjectName(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
			}
		})
	}
}

// TestResolveOutputDir tests output directory resolution
func TestResolveOutputDir(t *testing.T) {
	// Save and restore outputDir
	oldOutputDir := outputDir
	defer func() { outputDir = oldOutputDir }()

	t.Run("default to cwd", func(t *testing.T) {
		outputDir = ""
		cwd, _ := os.Getwd()

		dir, err := resolveOutputDir("myapp")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		expected := filepath.Join(cwd, "myapp")
		if dir != expected {
			t.Errorf("got %s, want %s", dir, expected)
		}
	})

	t.Run("custom output dir", func(t *testing.T) {
		outputDir = "/tmp/projects"

		dir, err := resolveOutputDir("myapp")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		// Should use the custom output dir as-is
		if dir != "/tmp/projects" {
			t.Errorf("got %s, want /tmp/projects", dir)
		}
	})
}

// TestCheckDirectory tests directory existence checking
func TestCheckDirectory(t *testing.T) {
	// Save and restore force flag
	oldForce := force
	defer func() { force = oldForce }()

	t.Run("non-existent directory", func(t *testing.T) {
		force = false
		err := checkDirectory("/path/that/does/not/exist")
		if err != nil {
			t.Errorf("expected no error for non-existent dir, got %v", err)
		}
	})

	t.Run("empty directory", func(t *testing.T) {
		force = false
		tmpDir, err := os.MkdirTemp("", "forge-test-*")
		if err != nil {
			t.Fatalf("failed to create temp dir: %v", err)
		}
		defer os.RemoveAll(tmpDir)

		err = checkDirectory(tmpDir)
		if err != nil {
			t.Errorf("expected no error for empty dir, got %v", err)
		}
	})

	t.Run("non-empty directory without force", func(t *testing.T) {
		force = false
		tmpDir, err := os.MkdirTemp("", "forge-test-*")
		if err != nil {
			t.Fatalf("failed to create temp dir: %v", err)
		}
		defer os.RemoveAll(tmpDir)

		// Create a file
		f, _ := os.Create(filepath.Join(tmpDir, "test.txt"))
		f.Close()

		err = checkDirectory(tmpDir)
		if err == nil {
			t.Error("expected error for non-empty dir without force")
		}
	})

	t.Run("non-empty directory with force", func(t *testing.T) {
		force = true
		tmpDir, err := os.MkdirTemp("", "forge-test-*")
		if err != nil {
			t.Fatalf("failed to create temp dir: %v", err)
		}
		defer os.RemoveAll(tmpDir)

		// Create a file
		f, _ := os.Create(filepath.Join(tmpDir, "test.txt"))
		f.Close()

		err = checkDirectory(tmpDir)
		if err != nil {
			t.Errorf("expected no error for non-empty dir with force, got %v", err)
		}
	})
}

// TestParseStackType tests stack type parsing
func TestParseStackType(t *testing.T) {
	tests := []struct {
		input string
		want  StackType
		ok    bool
	}{
		{"web", "web", true},
		{"mobile", "mobile", true},
		{"web-fullstack", "web-fullstack", true},
		{"invalid", "", false},
		{"", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got, ok := ParseStackType(tt.input)
			if ok != tt.ok {
				t.Errorf("ParseStackType(%q) ok = %v, want %v", tt.input, ok, tt.ok)
			}
			if ok && got != tt.want {
				t.Errorf("ParseStackType(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

// TestParseFeatures tests feature parsing
func TestParseFeatures(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{"empty", []string{}, 0},
		{"single", []string{"auth"}, 1},
		{"multiple", []string{"auth", "database", "api"}, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ParseFeatures(tt.input)
			if len(got) != tt.want {
				t.Errorf("ParseFeatures(%v) len = %d, want %d", tt.input, len(got), tt.want)
			}
		})
	}
}

// TestDirExists tests directory existence helper
func TestDirExists(t *testing.T) {
	t.Run("existing directory", func(t *testing.T) {
		tmpDir, _ := os.MkdirTemp("", "forge-test-*")
		defer os.RemoveAll(tmpDir)

		if !dirExists(tmpDir) {
			t.Error("expected true for existing directory")
		}
	})

	t.Run("non-existing directory", func(t *testing.T) {
		if dirExists("/path/that/does/not/exist") {
			t.Error("expected false for non-existing directory")
		}
	})

	t.Run("file not directory", func(t *testing.T) {
		tmpFile, _ := os.CreateTemp("", "forge-test-*")
		tmpFile.Close()
		defer os.Remove(tmpFile.Name())

		if dirExists(tmpFile.Name()) {
			t.Error("expected false for file")
		}
	})
}

// TestCommandExists tests command existence helper
func TestCommandExists(t *testing.T) {
	// These should exist on most systems
	if !commandExists("go") {
		t.Skip("go not found, skipping")
	}

	if commandExists("nonexistent-command-12345") {
		t.Error("expected false for non-existent command")
	}
}

// TestGetInvalidChars tests invalid character detection
func TestGetInvalidChars(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"myapp", ""},
		{"my-app", ""},
		{"my_app", ""},
		{"my app", " "},
		{"my@app", "@"},
		{"my@app!", "@!"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := GetInvalidChars(tt.input)
			if got != tt.want {
				t.Errorf("GetInvalidChars(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}
