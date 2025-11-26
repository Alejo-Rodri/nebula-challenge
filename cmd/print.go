package cmd

import (
/* 	"fmt"
	"os" */

	"github.com/Alejo-Rodri/nebula-challenge/internal/app"
	"github.com/Alejo-Rodri/nebula-challenge/internal/daemon"
	//"github.com/Alejo-Rodri/nebula-challenge/internal/infra/cli"
	"github.com/spf13/cobra"
)

func PrintCmd(app app.AssessmentStorage, socketPath string) *cobra.Command {
	var printCmd = &cobra.Command{
		Use: "print",
		Short: "Prints all the assessments done in the session",
		Long: `
		`,
		Run: func (cmd *cobra.Command, args []string)  {
			print(cmd, app, socketPath)
		},
	}

	printCmd.Flags().StringP("key", "k", "ssllabs", "Key used to search and print the results of the assessment")

	return printCmd
}

func print(cmd *cobra.Command, app app.AssessmentStorage, socketPath string) {

	daemon.ListValues(socketPath)
	/* assessmentKey, err := cmd.Flags().GetString("key")
	if err != nil {
		fmt.Fprintln(os.Stderr, HumanizeError(err))
	}

	result, err := app.Get(assessmentKey)
	if err != nil {
		fmt.Fprintln(os.Stderr, HumanizeError(err))
	}

	cli.PrintApiAnalyze(result) */
}
