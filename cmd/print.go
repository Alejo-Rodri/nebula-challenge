package cmd

import (
	"fmt"
	"os" 
	"log"

	"github.com/Alejo-Rodri/nebula-challenge/internal/app"
	"github.com/Alejo-Rodri/nebula-challenge/internal/daemon"

	"github.com/Alejo-Rodri/nebula-challenge/internal/infra/cli"
	"github.com/spf13/cobra"
)

func PrintCmd(app app.AssessmentStorage, unix *daemon.UnixClient) *cobra.Command {
	var printCmd = &cobra.Command{
		Use: "print",
		Short: "Prints all the assessments done in the session",
		Long: `
		`,
		Run: func (cmd *cobra.Command, args []string)  {
			print(cmd, app, unix)
		},
	}

	printCmd.Flags().StringP("key", "k", "ssllabs", "Key used to search and print the results of the assessment")

	return printCmd
}

func print(cmd *cobra.Command, app app.AssessmentStorage, unix *daemon.UnixClient) {

	data, err := unix.ListValues()
	if err != nil {
		fmt.Printf("ERRRR %w", err)
	}

	log.Print(data)

	assessmentKey, err := cmd.Flags().GetString("key")
	if err != nil {
		fmt.Fprintln(os.Stderr, HumanizeError(err))
	}

	result, err := app.GetByKey(assessmentKey)
	if err != nil {
		fmt.Fprintln(os.Stderr, HumanizeError(err))
	}

	cli.PrintApiAnalyze(result)
}
