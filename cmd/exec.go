package cmd

import (
	"obadiaspelembe/cves-guard/utils"
	"os"

	"github.com/spf13/cobra"
)

var execCommand = &cobra.Command{
	Use:   "exec",
	Short: "exec policy check into cves-report",
	Run: func(cmd *cobra.Command, args []string) {
		if !utils.ExecPolicy(cmd.Flag("policy").Value.String(), cmd.Flag("cves-report").Value.String()) {
			os.Exit(1)
		}
	},
}
