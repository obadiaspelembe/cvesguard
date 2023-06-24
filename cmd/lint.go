package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var lintCommand = &cobra.Command{
	Use:   "lint",
	Short: "lint policy and cves-report files",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Lint file: ", cmd.Flag("cves-report").Value.String())
	},
}
