package rules

import (
	"config_analyzer/model"
	"strings"
)

type AlgoritmRule struct{}

func (r AlgoritmRule) Check(key string, value interface{}) []model.Problem {
	v, ok := value.(string)
	if !ok {
		return nil
	}

	alg := strings.ToLower(v)

	if alg == "md5" || alg == "sha1" {
		return []model.Problem{
			{
				Severity:       model.HIGH,
				Message:        "weak cryptographic algorithm used: " + v,
				Recommendation: "use stronger algorithms like SHA-256 or bcrypt",
			},
		}
	}
	return nil
}
