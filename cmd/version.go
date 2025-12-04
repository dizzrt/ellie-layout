package cmd

import (
	"fmt"

	"github.com/dizzrt/ellie-layout/internal/conf"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of example",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(conf.Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
