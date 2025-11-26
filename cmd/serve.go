package cmd

import (
	"github.com/Alejo-Rodri/nebula-challenge/internal/daemon"
	"github.com/spf13/cobra"
)

func ServeCmd(socketPath string) *cobra.Command {
	var serveCmd = &cobra.Command{
		Use: "serve",
		Short: "Run daemon",
		Run: func (cmd *cobra.Command, args []string)  {
			daemon.RunServer(socketPath)
		},
	}

	return serveCmd
}
