package rules

import (
	"config_analyzer/model"
	"strings"
)

type DebugRule struct{}

func (r DebugRule) Check(key string, value interface{}) []model.Problem {
	if strings.ToLower(key) == "level" {
		if v, ok := value.(string); ok && strings.ToLower(v) == "debug" {
			return []model.Problem{
				{
					Severity:       model.LOW,
					Message:        "logging level is set to debug",
					Recommendation: "use info level or higher in production",
				},
			}
		}
	}
	return nil
}
