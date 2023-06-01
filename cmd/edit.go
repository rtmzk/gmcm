package cmd

import (
	"github.com/spf13/cobra"
	"gmcm/pkg/utils"
)

var DefaultConfigPath = utils.UserHome() + "/.env"

var edit = &cobra.Command{
	Use:   "edit",
	Short: "edit app config file",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := utils.OpenFileInEditor(DefaultConfigPath); err != nil {
			return err
		}
		return nil
	},
}
