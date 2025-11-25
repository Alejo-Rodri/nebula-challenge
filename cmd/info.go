package cmd

import (
	"fmt"
	"os"

	"github.com/Alejo-Rodri/nebula-challenge/internal/app"
	"github.com/spf13/cobra"
)

func InfoCmd(a app.AssessmentApp) *cobra.Command {
	return &cobra.Command{
		Use: "info",
		Short: "Shows information of the availability of the SSL labs server",
		Long: `
			This command should be used to check the availability of the SSL Labs servers, retrieve the engine and criteria version, and initialize the maximum number of concurrent assessments.
		`,
		Run: func (_ *cobra.Command, _ []string) {
			info(a)
		},
	}
}

func info(a app.AssessmentApp) {
	result, err := a.Info()
	if err != nil {
		// in this level the errors should be showed to the client of the app
		fmt.Fprintln(os.Stderr, HumanizeError(err))
		return
	}

	fmt.Printf("%+v", result)
}
