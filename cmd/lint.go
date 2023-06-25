package cmd

import (
	"obadiaspelembe/cves-guard/utils"

	"github.com/spf13/cobra"
)

var lintCommand = &cobra.Command{
	Use:   "lint",
	Short: "lint policy and cves-report files",
	Run: func(cmd *cobra.Command, args []string) {
		utils.Validate(cmd.Flag("policy").Value.String(), cmd.Flag("cves-report").Value.String())
	},
}
