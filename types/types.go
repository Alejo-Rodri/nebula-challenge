package types

type ApiErrorsResponse struct {
	Errors []struct {
		Field   string `json:"field"`
		Message string `json:"message"`
	} `json:"errors"`
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

