package model

type Severity string

const (
	LOW    Severity = "Low"
	MEDIUM Severity = "Medium"
	HIGH   Severity = "High"
)

type Problem struct {
	Severity       Severity
	Message        string
	Recommendation string
}
