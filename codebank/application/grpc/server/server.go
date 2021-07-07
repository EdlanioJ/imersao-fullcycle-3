package server

import (
	"fmt"
	"log"
	"net"

	"github.com/codeedu/codebank/application/grpc/pb"
	"github.com/codeedu/codebank/application/grpc/service"
	"github.com/codeedu/codebank/data/usecase"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GRPCServer struct {
	ProcessTransactionUseCase usecase.UseCaseTransaction
}

func NewGRPCServer() GRPCServer {
	return GRPCServer{}
}

func (g GRPCServer) Serve() {
	port := 50052

	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", port))
	if err != nil {
		log.Fatalf("could not listen tpc port")
	}

	transactionService := service.NewTransactionService()
	transactionService.ProcessTransactionUseCase = g.ProcessTransactionUseCase
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	pb.RegisterPaymentServiceServer(grpcServer, transactionService)

	log.Printf("gRPC server started at port %d", port)
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}
