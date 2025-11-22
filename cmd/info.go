package cmd

import (
	"github.com/Alejo-Rodri/nebula-challenge/configs"
	"github.com/Alejo-Rodri/nebula-challenge/internal/infra/api"
	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use: "info",
	Short: "Shows information of the availability of the SSL labs server",
	// TODO: Long description
	Run: info,
}

func init() {
	rootCmd.AddCommand(infoCmd)
}

func info(_ *cobra.Command, _ []string) {
	client := api.NewApiClient(configs.Envs.BaseApiURL)

	client.Info()
}