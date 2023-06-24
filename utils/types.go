package utils

const VERSION = "Version"
const SPEC = "Spec"
const NAME = "Name"

type SpecItem struct {
	Type   string `json:"type"`
	Target string `json:"target"`
}

type Policy struct {
	Version string     `json:"version"`
	Name    string     `json:"name"`
	Spec    []SpecItem `json:"spec"`
}

type ReportRunToolDriverRuleHelp struct {
	Text     string `json:"text"`
	Markdown string `json:"markdown"`
}

type ReportRunToolDriverRuleProperty struct {
	AffectedVersion   string   `json:"affected_version"`
	CvssV3Severity    string   `json:"cvss_severity"`
	FixedVersion      string   `json:"fixed_version"`
	PurlVersion       string   `json:"purl_version"`
	SecutirySerevirty bool     `json:"secutiry-serevirty"`
	Tags              []string `json:"tags"`
}
type ReportRunToolDriverRule struct {
	Id         string                          `json:"id"`
	Name       string                          `json:"name"`
	HelpUri    string                          `json:"helpUri"`
	Help       ReportRunToolDriverRuleHelp     `json:"help"`
	Properties ReportRunToolDriverRuleProperty `json:"properties"`
}

type ReportRunToolDriver struct {
	FullName       string                    `json:"fullName"`
	InformationUri string                    `json:"informationUri"`
	Name           string                    `json:"name"`
	Rules          []ReportRunToolDriverRule `json:"rules"`
	Version        string                    `json:"version"`
}
type ReportRunTool struct {
	Driver ReportRunToolDriver `json:"driver"`
}

type ReportRunResultMessage struct {
	Text string `json:"text"`
}
type ReportRunResult struct {
	RuleId    string                 `json:"ruleId"`
	RuleIndex string                 `json:"ruleIndex"`
	Kind      string                 `json:"kind"`
	Level     string                 `json:"level"`
	Message   ReportRunResultMessage `json:"message"`
}

type ReportRun struct {
	Tool    ReportRunTool     `json:"tool"`
	Results []ReportRunResult `json:"results"`
}

type Report struct {
	Version string            `json:"version"`
	Runs    []ReportRunResult `json:"runs"`
}
