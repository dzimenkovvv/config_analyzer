package rules

import (
	"config_analyzer/model"
	"strings"
)

type PermissionRule struct{}

func (r PermissionRule) Check(key string, value interface{}) []model.Problem {
	keyLower := strings.ToLower(key)

	if strings.Contains(keyLower, "perm") || strings.Contains(keyLower, "mode") || strings.Contains(keyLower, "chmode") {
		if v, ok := value.(string); ok {
			if v == "777" || v == "0666" {
				return []model.Problem{
					{
						Severity:       model.HIGH,
						Message:        "overly permissive access rights:" + v,
						Recommendation: "restrict permissions (e.g., 644 or 600)",
					},
				}
			}
		}
	}
	return nil
}
