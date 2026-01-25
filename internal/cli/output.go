package cli

import (
	"fmt"

	"github.com/fatih/color"
)

var (
	greenCheck = color.New(color.FgGreen).SprintFunc()
	redX       = color.New(color.FgRed).SprintFunc()
	blueInfo   = color.New(color.FgBlue).SprintFunc()
	yellowWarn = color.New(color.FgYellow).SprintFunc()
	bold       = color.New(color.Bold).SprintFunc()
)

// Success prints a success message with a green checkmark
func Success(msg string) {
	fmt.Printf("%s %s\n", greenCheck("✓"), msg)
}

// Error prints an error message with a red X
func Error(msg string) {
	fmt.Printf("%s %s\n", redX("✗"), msg)
}

// Info prints an info message with a blue arrow
func Info(msg string) {
	fmt.Printf("%s %s\n", blueInfo("→"), msg)
}

// Warn prints a warning message with a yellow warning sign
func Warn(msg string) {
	fmt.Printf("%s %s\n", yellowWarn("⚠"), msg)
}

// Header prints a bold header
func Header(msg string) {
	fmt.Printf("\n%s\n", bold(msg))
}

// PrintSummary prints the project creation summary with colors
func PrintSummary(projectName string, stack StackType, features []Feature, outputDir string) {
	Header("Project Summary")
	fmt.Printf("  Name:     %s\n", bold(projectName))
	fmt.Printf("  Stack:    %s\n", StackDisplayName(stack))
	fmt.Printf("  Features: %s\n", FormatFeatures(features))
	fmt.Printf("  Output:   %s\n", outputDir)
	fmt.Println()
}

// PrintNextSteps prints the next steps after project creation
func PrintNextSteps(projectName string, stack StackType, duration float64, postCreate PostCreateResult, skippedInstall bool) {
	Header("Next steps")
	fmt.Printf("  cd %s\n", projectName)

	// Show install steps only if not already done
	depsInstalled := postCreate.FrontendInstalled && postCreate.BackendInstalled

	// Stack-specific next steps based on stack ID
	switch string(stack) {
	case "web":
		if depsInstalled {
			fmt.Println("  npm run dev")
		} else if skippedInstall {
			fmt.Println("  npm install && npm run dev")
		} else {
			fmt.Println("  npm install && npm run dev")
		}
	case "mobile":
		if depsInstalled {
			fmt.Println("  npx expo start")
		} else if skippedInstall {
			fmt.Println("  npm install && npx expo start")
		} else {
			fmt.Println("  npm install && npx expo start")
		}
	}

	fmt.Println()

	// Show timing breakdown
	if duration > 0 {
		if depsInstalled && postCreate.Duration.Seconds() > 0 {
			fmt.Printf("Created in %.1fs (deps: %.1fs)\n", duration, postCreate.Duration.Seconds())
		} else {
			fmt.Printf("Created in %.1fs\n", duration)
		}
	}
}
