package repository

import (
	"sync"

	pb "shippy-service-consignment/proto/consignment"
)

type InMemoryRepository struct {
	mu           sync.RWMutex
	consignments []*pb.Consignment
}

func NewInMemoryRepository() InMemoryRepository {
	return InMemoryRepository{}
}

func (repo *InMemoryRepository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	repo.mu.Lock()

	updated := append(repo.consignments, consignment)
	repo.consignments = updated

	repo.mu.Unlock()

	return consignment, nil
}

func (repo *InMemoryRepository) GetAll() []*pb.Consignment {
	return repo.consignments
}
