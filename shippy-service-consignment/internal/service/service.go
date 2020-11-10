package service

import (
	"context"

	pb "shippy-service-consignment/proto/consignment"
)

// Необходимый интерфейс для сервисного слоя.
type ConsignmentRepository interface {
	Create(consignment *pb.Consignment) (*pb.Consignment, error)
	GetAll() []*pb.Consignment
}

type Consignment struct {
	repo ConsignmentRepository
}

// Сервис, реализующий бизнес логику.
func NewService(r ConsignmentRepository) Consignment {
	return Consignment{repo: r}
}

func (s *Consignment) CreateConsignment(_ context.Context, req *pb.Consignment) (*pb.Response, error) {
	consignment, err := s.repo.Create(req)
	if err != nil {
		return nil, err
	}

	return &pb.Response{Created: true, Consignment: consignment}, nil
}

func (s *Consignment) GetConsignments(_ context.Context, _ *pb.GetRequest) (*pb.Response, error) {
	consignments := s.repo.GetAll()
	return &pb.Response{Consignments: consignments}, nil
}
