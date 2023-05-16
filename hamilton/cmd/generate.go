package cmd

import (
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
