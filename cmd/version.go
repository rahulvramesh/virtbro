package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints the version of MyCLI",
	Long:  `All software has versions. This is MyCLI's version.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("virtbro v0.1.0")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
