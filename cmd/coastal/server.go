package main

import (
	"coastal/config/constant"
	"coastal/internal/app/coastal"
	"coastal/internal/pkg/pb"
	"fmt"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
)

func httpServer() {
	httpPort := constant.HttpPort

	server := coastal.New()

	s := grpc.NewServer()
	pb.RegisterImgtripServer(s, server)

	wrappedServer := grpcweb.WrapServer(s)
	handler := func(resp http.ResponseWriter, req *http.Request) {
		//resp.Header().Set("my-header", "*")
		wrappedServer.ServeHTTP(resp, req)
	}

	httpServer := http.Server{
		Addr:    fmt.Sprintf(":%d", httpPort),
		Handler: http.HandlerFunc(handler),
	}
	log.Printf("Starting server. http port: %d", httpPort)

	if err := httpServer.ListenAndServe(); err != nil {
		grpclog.Fatalf("failed starting http server: %v", err)
	}
}

func grpcServer() {
	rpcPort := fmt.Sprintf(":%d", constant.RpcPort)
	server := coastal.New()
	lis, err := net.Listen("tcp", rpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err.Error())
	}

	s := grpc.NewServer()
	pb.RegisterImgtripServer(s, server)
	reflection.Register(s)

	log.Printf("lestenning... %v", rpcPort)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
