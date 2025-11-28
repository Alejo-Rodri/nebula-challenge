package cmd

import (
	"github.com/Alejo-Rodri/nebula-challenge/internal/app"
	"github.com/Alejo-Rodri/nebula-challenge/internal/daemon"
	"github.com/spf13/cobra"
)

func ServeCmd(socketPath string, db app.AssessmentStorage) *cobra.Command {
	var serveCmd = &cobra.Command{
		Use: "serve",
		Short: "Run daemon",
		Run: func (cmd *cobra.Command, args []string)  {
			daemon.RunServer(socketPath, db)
		},
	}

	return serveCmd
}
