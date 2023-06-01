package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "gmcm",
	Short: "Used for install ceph rgw.",
}

func init() {
	rootCmd.AddCommand(edit)
	rootCmd.AddCommand(run)
	rootCmd.AddCommand(ver)
}

func Execute() error {
	return rootCmd.Execute()
}
