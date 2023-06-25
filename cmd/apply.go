package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"obadiaspelembe/cves-guard/utils"
)

var applyCommand = &cobra.Command{
	Use:   "apply",
	Short: "apply policy into cves-report",
	Run: func(cmd *cobra.Command, args []string) {
		if utils.ApplyPolicy(cmd.Flag("policy").Value.String(), cmd.Flag("cves-report").Value.String()) {
			fmt.Println("Configuration is valid")
		}
	},
}
