package cmd

import (
	"github.com/spf13/cobra"
	"gitlab.360live.vn/zpi/api/zpi-api-docs/service"
)

var runCommand = &cobra.Command{
	Short: "run command",
	Long:  `run command`,
	Run: func(cmd *cobra.Command, args []string) {
		service.InitService()
		service.Start()
	},
}
