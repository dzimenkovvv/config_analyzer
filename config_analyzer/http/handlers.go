package http

import (
	"config_analyzer/analyzer"
	"config_analyzer/parser"
	"encoding/json"
	"io"
	"net/http"
)

type HTTPHandlers struct {
	analyzer *analyzer.Analyzer
}

func NewHttpHandlers(a *analyzer.Analyzer) *HTTPHandlers {
	return &HTTPHandlers{
		analyzer: a,
	}
}

func (h *HTTPHandlers) HandleAnalyze(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	cfg, err := parser.Parse(data)
	if err != nil {
		http.Error(w, "parse error", http.StatusBadRequest)
		return
	}

	problems := h.analyzer.Analyze(cfg)

	w.Header()
	json.NewEncoder(w).Encode(problems)
}
