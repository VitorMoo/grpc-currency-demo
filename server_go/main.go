package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "demo/server/pb"
)

type server struct {
	pb.UnimplementedCurrencyServiceServer
}

func (s *server) GetRate(ctx context.Context, req *pb.RateRequest) (*pb.RateResponse, error) {
	log.Printf("Received rate request: from=%s, to=%s", req.FromCurrency, req.ToCurrency)

	var price float64
	if req.FromCurrency == "USD" && req.ToCurrency == "BRL" {
		price = 5.25
	} else {
		price = 0.0
	}

	return &pb.RateResponse{Price: price}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterCurrencyServiceServer(s, &server{})

	log.Printf("gRPC server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
