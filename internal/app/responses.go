package app

type Info struct {
    EngineVersion        string
    CriteriaVersion      string
    ClientMaxAssessments int
    MaxAssessments       int
    CurrentAssessments   int
    NewAssessmentCoolOff int
    Messages             []string
}

type Analysis struct {
    Host            string
    Port            int
    Protocol        string
    IsPublic        bool
    Status          string
    StartTime       int64
    TestTime        int64
    EngineVersion   string
    CriteriaVersion string
    Endpoints       []Endpoint
}

type Endpoint struct {
    IPAddress         string
    ServerName        string
    StatusMessage     string
    Grade             string
    GradeTrustIgnored string
    HasWarnings       bool
    IsExceptional     bool
    Progress          int
    Duration          int
    ETA               int
    Delegation        int
}
