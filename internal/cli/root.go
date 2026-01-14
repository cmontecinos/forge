package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "forge",
	Short: "Create projects with your preferred stacks",
	Long: `Forge is a CLI tool to create projects with your personalized stacks.

Available stacks:
  Web     Next.js (App Router + Tailwind + TypeScript) + Go (Echo) + Supabase
  Mobile  Expo (React Native + TypeScript) + Go (Echo) + Supabase

Optional features:
  auth      Login/registro via backend
  database  Conexión Go-Supabase
  api       Router, middlewares, handlers

Quick start:
  forge new my-app                              # Interactive mode
  forge new my-app --stack web --features auth  # Non-interactive

Each project is generated with your exact configurations, ready to start
building features immediately.`,
}

// Execute runs the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
