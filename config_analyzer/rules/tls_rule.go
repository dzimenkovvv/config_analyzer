package rules

import (
	"config_analyzer/model"
	"strings"
)

type TlsRule struct{}

func (r TlsRule) Check(key string, value interface{}) []model.Problem {
	keyLower := strings.ToLower(key)

	if strings.Contains(keyLower, "tls") {
		if v, ok := value.(bool); ok && !v {
			return []model.Problem{
				{
					Severity:       model.HIGH,
					Message:        "TLS is disabled",
					Recommendation: "enable TLS to secure communication",
				},
			}
		}
	}
	return nil
}
