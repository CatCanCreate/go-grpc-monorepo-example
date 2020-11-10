package grpchandler

import (
	"context"

	"shippy-service-consignment/internal/service"
	pb "shippy-service-consignment/proto/consignment"
)

type Server struct {
	pb.UnimplementedShippingServiceServer
	service service.Consignment
}

// Сервис, реализующий gRPC интерфейс.
func NewServer(s service.Consignment) Server {
	return Server{service: s}
}

func (s *Server) CreateConsignment(ctx context.Context, req *pb.Consignment) (*pb.Response, error) {
	consignment, err := s.service.CreateConsignment(ctx, req)
	return consignment, err
}

func (s *Server) GetConsignments(ctx context.Context, req *pb.GetRequest) (*pb.Response, error) {
	consignments, err := s.service.GetConsignments(ctx, req)
	return consignments, err
}
