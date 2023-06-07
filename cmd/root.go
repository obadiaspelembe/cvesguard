package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
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
	if err := rootCommand.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
