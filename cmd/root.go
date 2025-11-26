/*
Copyright © 2025 Alejandro Rodriguez alejandrorodriguezq0@gmail.com
*/
package cmd

import (
	"os"

	"github.com/Alejo-Rodri/nebula-challenge/configs"
	"github.com/Alejo-Rodri/nebula-challenge/internal/infra/api"
	"github.com/Alejo-Rodri/nebula-challenge/internal/infra/db"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "nebula-challenge",
	Short: "Checks the TLS security of a given domain",
	Long: `
		Nebula Challenge is a CLI tool written in Go that checks the TLS security
		of a given domain using the SSL Labs API.

		It allows you to:
		- Retrieve general information about the SSL Labs servers
		- Analyze a domain or IP and track its assessment status
		- View detailed TLS results once the analysis is ready

		This tool is useful for quick checks, scripting, automation,
		and understanding how a domain scores in SSL Labs without
		opening the web interface.
	`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	injectDeps()
	
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var socketPath string

func injectDeps() {
    client := api.NewApiClient(configs.Envs.BaseApiURL)
	infoReq := api.NewInfoRequest(client)
	analyzeReq := api.NewAnalyzeRequest(client)
	assManager := db.NewAssessmentManager()

    rootCmd.AddCommand(InfoCmd(client, infoReq))
    rootCmd.AddCommand(AnalyzeCmd(client, analyzeReq, &assManager, socketPath))
	rootCmd.AddCommand(PrintCmd(&assManager, socketPath))
	rootCmd.AddCommand(ServeCmd(socketPath))

	rootCmd.PersistentFlags().StringVar(&socketPath, "socket", "/tmp/nebula-challenge.sock", "unix socket path")
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.nebula-challenge.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

