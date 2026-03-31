package scanner

import (
	"config_analyzer/analyzer"
	"config_analyzer/model"
	"config_analyzer/parser"
	"io/fs"
	"os"
	"path/filepath"
)

func ScanPath(path string, a *analyzer.Analyzer) ([]model.Problem, error) {
	var problems []model.Problem

	info, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	if !info.IsDir() {
		return scanFile(path, a)
	}

	err = filepath.WalkDir(path, func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}

		if d.IsDir() {
			return nil
		}

		fileProblems, err := scanFile(p, a)
		if err == nil {
			problems = append(problems, fileProblems...)
		}
		return nil
	})

	return problems, err
}

func scanFile(path string, a *analyzer.Analyzer) ([]model.Problem, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	cfg, err := parser.Parse(data)
	if err != nil {
		return nil, err
	}

	return a.Analyze(cfg), nil
}
