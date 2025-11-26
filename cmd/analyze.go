package cmd

import (
	"fmt"
	"os"

	"github.com/Alejo-Rodri/nebula-challenge/internal/app"
	"github.com/Alejo-Rodri/nebula-challenge/internal/daemon"
	//"github.com/Alejo-Rodri/nebula-challenge/internal/infra/cli"
	"github.com/spf13/cobra"
)

func AnalyzeCmd(
	a app.AssessmentApp,
	get app.GetRequest[app.Analysis],
	storage app.AssessmentStorage,
	socketPath string,
) *cobra.Command {
	var analyzeCmd = &cobra.Command{
		Use: "analyze",
		Short: "Analyzes a domain or ip address and stores its value",
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
			analyze(cmd, a, get, storage, socketPath)
		},
	}

	analyzeCmd.Flags().StringP("domain", "d", "www.ssllabs.com", "Domain or ip address to analyze")

	analyzeCmd.Flags().BoolP("process", "p", false, "Indicates whether the command should run in the background")
	analyzeCmd.Flags().StringP("key", "k", "", "Key used to save the results of the assessment, if empty the key is the url")

	return analyzeCmd
}

func analyze(
	cmd *cobra.Command, a app.AssessmentApp,
	get app.GetRequest[app.Analysis],
	storage app.AssessmentStorage,
	socketPath string,
) {
	host, err := cmd.Flags().GetString("domain")
	if err != nil {
		fmt.Fprintln(os.Stderr, HumanizeError(err))
	}

	/* isProcess, err := cmd.Flags().GetBool("process")
	if err != nil {
		fmt.Fprintln(os.Stderr, HumanizeError(err))
	} */

	assessmentKey, err := cmd.Flags().GetString("key")
	if err != nil {
		fmt.Fprintln(os.Stderr, HumanizeError(err))
	}

	if assessmentKey == "" {
		assessmentKey = host
	}

	/* result, err := a.Analyze(host, isProcess, get)
	if err != nil {
		// in this level the errors should be showed to the client of the app
		fmt.Fprintln(os.Stderr, HumanizeError(err))
		return
	} */

	// se le inyectaria la funcion para almacenar el resultado y aca se llamaria
	//storage.Save(assessmentKey, result)
	daemon.AddValue(socketPath, assessmentKey)

	//cli.PrintApiAnalyze(result)
}
