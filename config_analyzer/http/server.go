package http

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

type HTTPServer struct {
	httpHandlers *HTTPHandlers
}

func NewHttpServer(h *HTTPHandlers) *HTTPServer {
	return &HTTPServer{
		httpHandlers: h,
	}
}

func (s *HTTPServer) StartServer() error {
	router := mux.NewRouter()

	router.HandleFunc("/analyze", s.httpHandlers.HandleAnalyze).Methods("POST")

	if err := http.ListenAndServe(":9091", router); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return err
	}
	return nil
}
