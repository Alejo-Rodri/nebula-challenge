package cmd

import (
	"fmt"
	"os" 

	"github.com/Alejo-Rodri/nebula-challenge/internal/daemon"

	"github.com/Alejo-Rodri/nebula-challenge/internal/infra/cli"
	"github.com/spf13/cobra"
)

func PrintCmd(unix *daemon.UnixClient) *cobra.Command {
	var printCmd = &cobra.Command{
		Use: "print",
		Short: "Prints all the assessments done in the session",
		Long: `
		The print command displays the assessments stored in the current session.

		If no key is provided, it prints all available assessments and updates all not ready assessments.
		If a key is provided, it prints the detailed information of the matching assessment.

		Examples:
		nebula-challenge print
		nebula-challenge print -k my-key
		`,
		Run: func (cmd *cobra.Command, args []string)  {
			print(cmd, unix)
		},
	}

	printCmd.Flags().StringP("key", "k", "", "Key used to search and print the results of the assessment")

	return printCmd
}

func print(cmd *cobra.Command, unix *daemon.UnixClient) {
	assessmentKey, err := cmd.Flags().GetString("key")
	if err != nil {
		fmt.Fprintln(os.Stderr, HumanizeError(err))
	}

	if assessmentKey == "" {
		results, err := unix.ListAllValues()
		if err != nil {
			fmt.Fprintln(os.Stderr, HumanizeError(err))
		}

		cli.PrintAllResults(results)
		return
	}

	result, err := unix.GetAssResultByKey(assessmentKey)
	if err != nil {
		fmt.Fprintln(os.Stderr, HumanizeError(err))
	}

	cli.PrintApiAnalyze(result)
}
