package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"shippy-service-consignment/internal/grpchandler"
	"shippy-service-consignment/internal/repository"
	"shippy-service-consignment/internal/service"
	pb "shippy-service-consignment/proto/consignment"
)

const (
	port = ":50051"
)

func main() {
	inMemoryRepository := repository.NewInMemoryRepository()
	shippingService := service.NewService(&inMemoryRepository)
	shippingServer := grpchandler.NewServer(shippingService)

	//nolint:gosec
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln(err)
	}

	s := grpc.NewServer()

	pb.RegisterShippingServiceServer(s, &shippingServer)

	log.Println("Running on port:", port)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
