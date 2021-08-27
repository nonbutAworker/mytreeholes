package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of TreeHoles",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Personnel TreeHoles by zz v0.9 -- HEAD")
	},
}
