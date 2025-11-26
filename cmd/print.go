package cmd

import "github.com/spf13/cobra"

func PrintCmd() *cobra.Command {
	var printCmd = &cobra.Command{
		Use: "print",
		Short: "Prints all the assessments done in the session",
		Long: `
		`,
		Run: func (cmd *cobra.Command, args []string)  {
			print(cmd)		
		},
	}

	printCmd.Flags().StringP("key", "k", "ssllabs", "Key used to search and print the results of the assessment")

	return printCmd
}

func print(cmd *cobra.Command) {

}