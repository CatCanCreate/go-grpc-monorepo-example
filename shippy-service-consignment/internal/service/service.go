package service

import (
	"context"

	pb "shippy-service-consignment/proto/consignment"
)

// Необходимый интерфейс для сервисного слоя.
type Repository interface {
	Create(consignment *pb.Consignment) (*pb.Consignment, error)
	GetAll() []*pb.Consignment
}

type Service struct {
	repo Repository
}

// Сервис, реализующий gRPC интерфейс.
func NewService(r Repository) Service {
	return Service{repo: r}
}

func (s *Service) CreateConsignment(_ context.Context, req *pb.Consignment) (*pb.Response, error) {
	consignment, err := s.repo.Create(req)
	if err != nil {
		return nil, err
	}

	return &pb.Response{Created: true, Consignment: consignment}, nil
}

func (s *Service) GetConsignments(_ context.Context, _ *pb.GetRequest) (*pb.Response, error) {
	consignments := s.repo.GetAll()
	return &pb.Response{Consignments: consignments}, nil
}
