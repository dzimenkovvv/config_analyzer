package main

import (
	"config_analyzer/analyzer"
	"config_analyzer/config_analyzer/grpc/pb"
	"config_analyzer/grpc/server"
	"config_analyzer/rules"
	"log"
	"net"

	"google.golang.org/grpc"
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

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()

	pb.RegisterAnalyzeServiceServer(s, server.NewServer(a))

	log.Println("gRPC server running on: :50051")

	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
