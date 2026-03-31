package main

import (
	"config_analyzer/analyzer"
	"config_analyzer/model"
	"config_analyzer/parser"
	"config_analyzer/rules"
	"config_analyzer/scanner"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	silent := flag.Bool("s", false, "silent mode")
	stdin := flag.Bool("stdin", false, "read from stdin")

	flag.Parse()

	a := analyzer.NewAnalyzer([]rules.Rule{
		rules.AlgoritmRule{},
		rules.DebugRule{},
		rules.HostRule{},
		rules.PasswordRule{},
		rules.PermissionRule{},
		rules.TlsRule{},
	})

	var problems []model.Problem
	var err error

	if *stdin {
		data, err := io.ReadAll(os.Stdin)
		if err != nil {
			fmt.Println("error readin stdin:", err)
			os.Exit(1)
		}

		cfg, err := parser.Parse(data)
		if err != nil {
			fmt.Println("parse error:", err)
			os.Exit(1)
		}

		problems = a.Analyze(cfg)
	} else {
		if flag.NArg() == 0 {
			fmt.Println("please provide config file or directory path")
			os.Exit(1)
		}

		path := flag.Arg(0)

		problems, err = scanner.ScanPath(path, a)
		if err != nil {
			fmt.Println("scan error:", err)
			os.Exit(1)
		}
	}

	if len(problems) > 0 {
		for _, p := range problems {
			fmt.Printf("[%s] %s\n%s\n\n", p.Severity, p.Message, p.Recommendation)
		}

		if !*silent {
			os.Exit(1)
		}
	}
}
