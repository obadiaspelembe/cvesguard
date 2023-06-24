package utils

const VERSION = "Version"
const SPEC = "Spec"
const NAME = "Name"

type SpecItem struct {
	Type   string `json:"type" binding:"required"`
	Target string `json:"target" binding:"required"`
}

type Policy struct {
	Version string     `json:"version"`
	Name    string     `json:"name"`
	Spec    []SpecItem `json:"spec"`
}


