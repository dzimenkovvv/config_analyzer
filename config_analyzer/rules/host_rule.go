package rules

import "config_analyzer/model"

type HostRule struct{}

func (r HostRule) Check(key string, value interface{}) []model.Problem {
	if v, ok := value.(string); ok && v == "0.0.0.0" {
		return []model.Problem{
			{
				Severity:       model.MEDIUM,
				Message:        "application is bound to 0.0.0.0",
				Recommendation: "bind to a specific interface or restrict access",
			},
		}
	}
	return nil
}
