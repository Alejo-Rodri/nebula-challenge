package cmd

import (
	"fmt"
	"os"

	"github.com/Alejo-Rodri/nebula-challenge/internal/app"
	"github.com/spf13/cobra"
)

func AnalyzeCmd(a app.AssessmentApp) *cobra.Command {
	var analyzeCmd = &cobra.Command{
		Use: "analyze",
		Short: "Analyzes a domain or ip address",
		// TODO add description of cmd
		Long: `
			
		`,
		Run: func (cmd *cobra.Command, _ []string)  {
			analyze(cmd, a)
		},
	}

	analyzeCmd.Flags().StringP("domain", "d", "www.ssllabs.com", "Domain or ip address to analyze")

	return analyzeCmd
}

func analyze(cmd *cobra.Command, a app.AssessmentApp) {
	host, err := cmd.Flags().GetString("domain")
	if err != nil {
		fmt.Fprintln(os.Stderr, HumanizeError(err))
	}

	result, err := a.Analyze(host)
	if err != nil {
		// in this level the errors should be showed to the client of the app
		fmt.Fprintln(os.Stderr, HumanizeError(err))
		return
	}

	fmt.Printf("%+v", result)
}