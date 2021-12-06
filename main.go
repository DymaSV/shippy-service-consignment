package main

import (
	"sync"

	pb "github.com/DymaSV/shippy-service-consignment/shippy-service-consignment/proto/consignment"
)

const (
	port = ":50051"
)

type Repository struct {
	mu           sync.RWMutex
	conseignment []*pb.Consignment
}

func (repo *Repository) Create(conseignment *pb.Consignment) (*pb.Consignment, error) {

}
