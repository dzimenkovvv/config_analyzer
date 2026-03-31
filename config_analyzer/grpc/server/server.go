package server

import (
	"config_analyzer/analyzer"
	"config_analyzer/config_analyzer/grpc/pb"
	"config_analyzer/parser"
	"context"
)

type Server struct {
	pb.UnimplementedAnalyzeServiceServer
	analyzer *analyzer.Analyzer
}

func NewServer(a *analyzer.Analyzer) *Server {
	return &Server{
		analyzer: a,
	}
}

func (s *Server) Analyze(ctx context.Context, req *pb.AnalyzeRequest) (*pb.AnalyzeResponse, error) {
	cfg, err := parser.Parse(req.Config)
	if err != nil {
		return nil, err
	}

	problems := s.analyzer.Analyze(cfg)

	resp := &pb.AnalyzeResponse{}

	for _, p := range problems {
		resp.Problems = append(resp.Problems, &pb.Problem{
			Severity:       string(p.Severity),
			Message:        p.Message,
			Recommendation: p.Recommendation,
		})
	}
	return resp, nil
}
