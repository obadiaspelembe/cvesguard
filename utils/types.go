package utils

type Configuration struct {
	Critical *int `json:"critical" binding:"required"`
	High     *int `json:"high" binding:"required"`
	Medium   *int `json:"medium" binding:"required"`
	Low      *int `json:"low" binding:"required"`
}

type Spec struct {
	Config *Configuration `json:"config" binding:"required"`
}

type Policy struct {
	Version *string `json:"version" binding:"required"`
	Name    *string `json:"name" binding:"required"`
	Kind    *string `json:"kind" binding:"required"`
	Spec    *Spec   `json:"spec" binding:"required"`
}

type ReportRunToolDriverRuleHelp struct {
	Text     *string `json:"text" binding:"required"`
	Markdown *string `json:"markdown" binding:"required"`
}

type ReportRunToolDriverRuleProperty struct {
	Affected_Version *string   `json:"affected_version" binding:"required"`
	CvssV3_Severity  *string   `json:"cvssV3_severity" binding:"required"`
	Fixed_Version    *string   `json:"fixed_version" binding:"required"`
	Purl             *string   `json:"purl" binding:"required"`
	SecuritySeverity *string   `json:"security-severity" binding:"required"`
	Tags             *[]string `json:"tags" binding:"required"`
}
type ReportRunToolDriverRule struct {
	Id               *string                          `json:"id" binding:"required"`
	Name             *string                          `json:"name" binding:"required"`
	HelpUri          *string                          `json:"helpUri" binding:"required"`
	ShortDescription *ReportRunResultMessage          `json:"shortDescription" binding:"required"`
	Help             *ReportRunToolDriverRuleHelp     `json:"help" binding:"required"`
	Properties       *ReportRunToolDriverRuleProperty `json:"properties" binding:"required"`
}

type ReportRunToolDriver struct {
	FullName       *string                    `json:"fullName" binding:"required"`
	InformationUri *string                    `json:"informationUri" binding:"required"`
	Name           *string                    `json:"name" binding:"required"`
	Rules          *[]ReportRunToolDriverRule `json:"rules" binding:"required"`
	Version        *string                    `json:"version"`
}
type ReportRunTool struct {
	Driver *ReportRunToolDriver `json:"driver" binding:"required"`
}

type ReportRunResultMessage struct {
	Text *string `json:"text" binding:"required"`
}
type ReportRunResult struct {
	RuleId    *string                 `json:"ruleId" binding:"required"`
	RuleIndex *string                 `json:"ruleIndex" binding:"required"`
	Kind      *string                 `json:"kind" binding:"required"`
	Level     *string                 `json:"level" binding:"required"`
	Message   *ReportRunResultMessage `json:"message" binding:"required"`
}

type ReportRun struct {
	Tool    *ReportRunTool     `json:"tool" binding:"required"`
	Results *[]ReportRunResult `json:"results" binding:"required"`
}

type Report struct {
	Version *string      `json:"version" binding:"required"`
	Runs    *[]ReportRun `json:"runs" binding:"required"`
}
