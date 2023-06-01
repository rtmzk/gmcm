package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gmcm/pkg/utils/version"
	yaml "gopkg.in/yaml.v3"
)

var ver = &cobra.Command{
	Use:   "version",
	Short: "show server version",
	Run: func(cmd *cobra.Command, args []string) {
		v := version.Get()
		marshaled, err := yaml.Marshal(&v)
		if err != nil {
			return
		}
		fmt.Println(string(marshaled))
	},
}
