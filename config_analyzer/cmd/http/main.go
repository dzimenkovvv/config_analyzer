package main

import (
	"config_analyzer/analyzer"
	"config_analyzer/http"
	"config_analyzer/rules"
)

func main() {
	a := analyzer.NewAnalyzer([]rules.Rule{
		rules.AlgoritmRule{},
		rules.DebugRule{},
		rules.HostRule{},
		rules.PasswordRule{},
		rules.PermissionRule{},
		rules.TlsRule{},
	})

	h := http.NewHttpHandlers(a)

	s := http.NewHttpServer(h)

	if err := s.StartServer(); err != nil {
		panic(err)
	}
}
