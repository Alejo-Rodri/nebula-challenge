package api

type ApiErrorsResponse struct {
	Errors []ApiError `json:"errors"`
}

type ApiError struct {
    Field   string `json:"field"`
    Message string `json:"message"`
}

type ApiInfoResponse struct {
	EngineVersion        string   `json:"engineVersion"`
	CriteriaVersion      string   `json:"criteriaVersion"`
	ClientMaxAssessments int      `json:"clientMaxAssessments"`
	MaxAssessments       int      `json:"maxAssessments"`
	CurrentAssessments   int      `json:"currentAssessments"`
	NewAssessmentCoolOff int      `json:"newAssessmentCoolOff"`
	Messages             []string `json:"messages"`
}

type ApiAnalyzeResponse struct {
	Host            string `json:"host"`
	Port            int    `json:"port"`
	Protocol        string `json:"protocol"`
	IsPublic        bool   `json:"isPublic"`
	Status          string `json:"status"`
	StartTime       int64  `json:"startTime"`
	TestTime        int64  `json:"testTime"`
	EngineVersion   string `json:"engineVersion"`
	CriteriaVersion string `json:"criteriaVersion"`
	Endpoints       []ApiAnalyzeEndpoints `json:"endpoints"`
}

type ApiAnalyzeEndpoints struct {
	IPAddress         string `json:"ipAddress"`
	ServerName        string `json:"serverName"`
	StatusMessage     string `json:"statusMessage"`
	Grade             string `json:"grade"`
	GradeTrustIgnored string `json:"gradeTrustIgnored"`
	HasWarnings       bool   `json:"hasWarnings"`
	IsExceptional     bool   `json:"isExceptional"`
	Progress          int    `json:"progress"`
	Duration          int    `json:"duration"`
	Eta               int    `json:"eta"`
	Delegation        int    `json:"delegation"`
}
