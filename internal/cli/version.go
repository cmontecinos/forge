package cli

import (
	"fmt"

	"github.com/bigbytes/forge/internal/config"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version information",
	Long:  `Display the version, build date, and git commit of forge.`,
	Run: func(cmd *cobra.Command, args []string) {
		printVersion()
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)

	// Add --version flag to root command
	rootCmd.Version = config.Version
	rootCmd.SetVersionTemplate(fmt.Sprintf("forge version %s\n", config.Version))
}

func printVersion() {
	fmt.Printf("forge version %s (built: %s, commit: %s)\n",
		config.Version,
		config.BuildDate,
		config.GitCommit,
	)
}
