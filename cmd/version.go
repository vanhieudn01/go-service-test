package cmd

import (
	"expvar"
	"fmt"

	"github.com/spf13/cobra"
)

var versionCommand = &cobra.Command{
	Short: "version command",
	Long:  `version command`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("service-test version", version)
	},
}

var version = "0.0.1"

//SetRevision ...
func SetRevision(r string) {
	if len(r) > 0 {
		version = fmt.Sprintf("%v-%v", version, r)
	}

	expvar.NewString("service-test-version").Set(version)
}
