package cmd

import (
	"github.com/Alejo-Rodri/nebula-challenge/internal/app"
	"github.com/Alejo-Rodri/nebula-challenge/internal/daemon"
	"github.com/spf13/cobra"
)

func ServeCmd(socketPath string, db app.AssessmentStorage, analyzer app.Analize) *cobra.Command {
	var serveCmd = &cobra.Command{
		Use: "serve",
		Short: "Run daemon",
		Long: `
			The serve command starts the daemon responsible for storing and managing
			assessments. It runs a background server that communicates with the CLI through
			Unix domain sockets.

			This daemon handles all read and write operations for assessment data, allowing
			other commands to request, update, or retrieve results without blocking.

			Example:
			nebula-challenge serve
		`,
		Run: func (cmd *cobra.Command, args []string)  {
			daemon.RunServer(socketPath, db, analyzer)
		},
	}

	return serveCmd
}
