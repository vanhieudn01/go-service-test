package cmd

import (
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
	"go.uber.org/zap"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

func initCmd() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigName("config")
	}

	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetConfigType("toml")

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("viper using config file", viper.ConfigFileUsed())
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(fs fsnotify.Event) {
		fmt.Println("config file changed", zap.Any("event", fs))
	})
}

func init() {
	cobra.OnInitialize(initCmd)
}

//RootCmd ...
var RootCmd = &cobra.Command{
	Short: "service test for common case",
	Long:  `service test for common case`,
	Use:   "service-test",
}

//Execute ...
func Execute() {
	RootCmd.AddCommand(runCommand)
	RootCmd.AddCommand(versionCommand)

	if err := RootCmd.Execute(); err != nil {
		log.Fatal("start command failed", zap.Error(err))
	}
}
