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
		`,
		Run: func (cmd *cobra.Command, args []string)  {
			print(cmd, unix)
		},
	}

	printCmd.Flags().StringP("key", "k", "", "Key used to search and print the results of the assessment")

	return printCmd
}

// TODO improve errors
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
	}

	result, err := unix.GetAssResultByKey(assessmentKey)
	if err != nil {
		fmt.Fprintln(os.Stderr, HumanizeError(err))
	}

	cli.PrintApiAnalyze(result)
}
