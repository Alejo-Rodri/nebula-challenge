package cli

import (
	"fmt"

	"github.com/Alejo-Rodri/nebula-challenge/internal/infra/api"
)

func PrintApiInfo(info api.ApiInfoResponse) {
	fmt.Printf(
		"Engine Version:        %s\n"+
			"Criteria Version:      %s\n"+
			"Client Max Assessments: %d\n"+
			"Max Assessments:        %d\n"+
			"Current Assessments:    %d\n"+
			"New Assessment CoolOff: %d\n"+
			"Messages:\n",
		info.EngineVersion,
		info.CriteriaVersion,
		info.ClientMaxAssessments,
		info.MaxAssessments,
		info.CurrentAssessments,
		info.NewAssessmentCoolOff,
	)

	for _, msg := range info.Messages {
		fmt.Printf("  - %s\n", msg)
	}
}

func PrintApiAnalyze(resp api.ApiAnalyzeResponse) {
	fmt.Printf(
		"Host:              %s\n"+
			"Port:              %d\n"+
			"Protocol:          %s\n"+
			"Public:            %t\n"+
			"Status:            %s\n"+
			"Engine Version:    %s\n"+
			"Criteria Version:  %s\n"+
			"\nEndpoints:\n",
		resp.Host,
		resp.Port,
		resp.Protocol,
		resp.IsPublic,
		resp.Status,
		resp.EngineVersion,
		resp.CriteriaVersion,
	)

	for i, ep := range resp.Endpoints {
		fmt.Printf("  [%d]\n", i+1)
		fmt.Printf("    IP:             %s\n", ep.IPAddress)
		fmt.Printf("    Server Name:    %s\n", ep.ServerName)
		fmt.Printf("    Status Message: %s\n", ep.StatusMessage)

		if ep.Grade != "" {
			fmt.Printf("    Grade:          %s\n", ep.Grade)
			fmt.Printf("    Grade (Trust):  %s\n", ep.GradeTrustIgnored)
		}

		if ep.HasWarnings {
			fmt.Printf("    Warnings:       yes\n")
		} else {
			fmt.Printf("    Warnings:       no\n")
		}

		if ep.IsExceptional {
			fmt.Printf("    Exceptional:    yes\n")
		} else {
			fmt.Printf("    Exceptional:    no\n")
		}

		fmt.Printf("    Progress:       %d%%\n", ep.Progress)
		fmt.Printf("    ETA:            %ds\n", ep.Eta)
		fmt.Printf("    Duration:       %dms\n", ep.Duration)
		fmt.Printf("    Delegation:     %d\n", ep.Delegation)

		fmt.Println()
	}
}