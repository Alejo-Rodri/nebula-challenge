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
	
}
