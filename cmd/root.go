package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

const SHORT_DESCRIPTION = "cvesguard is a policy management tool to evaluate docker scout command output."

var version = "1.0.0-alpha"
var rootCommand = &cobra.Command{
	Use:     "cvesguard",
	Version: version,
	Short:   SHORT_DESCRIPTION,
	Long:    SHORT_DESCRIPTION,
}

func Execute() {

	policy := ""
	report := ""

	lintCommand.Flags().StringVarP(&policy, "policy", "p", "", "Policy is required")
	lintCommand.Flags().StringVarP(&report, "cves-report", "r", "", "CVES report filename")
	
	applyCommand.Flags().StringVarP(&policy, "policy", "p", "", "Policy is required")
	applyCommand.Flags().StringVarP(&report, "cves-report", "r", "", "CVES report filename")
	
	lintCommand.MarkFlagRequired("cves-report")
	applyCommand.MarkFlagRequired("cves-report")


	rootCommand.AddCommand(lintCommand)
	rootCommand.AddCommand(applyCommand)

	if err := rootCommand.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
