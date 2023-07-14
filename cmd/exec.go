package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"obadiaspelembe/cves-guard/utils"
)

var execCommand = &cobra.Command{
	Use:   "exec",
	Short: "exec policy check into cves-report",
	Run: func(cmd *cobra.Command, args []string) {
		if utils.ExecPolicy(cmd.Flag("policy").Value.String(), cmd.Flag("cves-report").Value.String()) {
			fmt.Println("Configuration is valid")
		}
	},
}
