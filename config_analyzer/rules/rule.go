package rules

import "config_analyzer/model"

type Rule interface {
	Check(key string, value interface{}) []model.Problem
}
