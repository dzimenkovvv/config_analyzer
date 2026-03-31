package analyzer

import (
	"config_analyzer/model"
	"config_analyzer/rules"
)

type Analyzer struct {
	rules []rules.Rule
}

func NewAnalyzer(r []rules.Rule) *Analyzer {
	return &Analyzer{
		rules: r,
	}
}

func (a Analyzer) Analyze(cfg map[string]interface{}) []model.Problem {
	var problems []model.Problem

	var walk func(map[string]interface{})

	walk = func(m map[string]interface{}) {
		for k, v := range m {
			for _, rule := range a.rules {
				problems = append(problems, rule.Check(k, v)...)
			}

			if nested, ok := v.(map[string]interface{}); ok {
				walk(nested)
			}
		}
	}

	walk(cfg)
	return problems
}
