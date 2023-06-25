package cmd

import (
	"obadiaspelembe/cves-guard/utils"

	"github.com/spf13/cobra"
)

var applyCommand = &cobra.Command{
	Use:   "apply",
	Short: "apply policy into cves-report",
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("apply file: ", cmd.Flag("cves-report").Value.String())

		utils.ApplyPolicy(cmd.Flag("policy").Value.String(), cmd.Flag("cves-report").Value.String())
	},
}
