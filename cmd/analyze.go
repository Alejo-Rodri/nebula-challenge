package cmd

import (
	"fmt"
	"os"

	"github.com/Alejo-Rodri/nebula-challenge/internal/app"
	"github.com/Alejo-Rodri/nebula-challenge/internal/infra/api"
	"github.com/Alejo-Rodri/nebula-challenge/internal/infra/cli"
	"github.com/spf13/cobra"
)

func AnalyzeCmd(a app.AssessmentApp) *cobra.Command {
	var analyzeCmd = &cobra.Command{
		Use: "analyze",
		Short: "Analyzes a domain or ip address",
		Long: `
			The analyze command sends a request to the SSL Labs API to evaluate the TLS
			security of a given domain or IP address.

			It starts an assessment (or retrieves an existing one) and waits through the
			different analysis states until the final result is ready or an error occurs.

			Usage examples:
			nebula-challenge analyze -d example.com
			nebula-challenge analyze --domain example.com

			The command prints the assessment status, server grades, and other relevant
			TLS details once the analysis is complete.
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

	result, err := a.Analyze(host, api.Get[api.ApiAnalyzeResponse])
	if err != nil {
		// in this level the errors should be showed to the client of the app
		fmt.Fprintln(os.Stderr, HumanizeError(err))
		return
	}

	cli.PrintApiAnalyze(result)
}
