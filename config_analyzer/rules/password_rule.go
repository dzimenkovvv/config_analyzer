package rules

import (
	"config_analyzer/model"
	"strings"
)

type PasswordRule struct{}

func (r PasswordRule) Check(key string, value interface{}) []model.Problem {
	if strings.Contains(strings.ToLower(key), "password") {
		if v, ok := value.(string); ok && v != "" {
			return []model.Problem{
				{
					Severity:       model.HIGH,
					Message:        "password in plain text",
					Recommendation: "store passwords in environment variables or secure storage",
				},
			}
		}
	}
	return nil
}
