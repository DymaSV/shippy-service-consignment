package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"

	pb "github.com/DymaSV/shippy-service-consignment/proto/consignment"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

type Repository struct {
	mu           sync.RWMutex
	conseignment []*pb.Consignment
}

// Create a new consignment
func (repo *Repository) Create(conseignment *pb.Consignment) (*pb.Consignment, error) {
	repo.mu.Lock()
	repo.conseignment = append(repo.conseignment, conseignment)
	repo.mu.Unlock()
	return conseignment, nil
}

type service struct {
	repo *Repository
}

// Create methode for our service
func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment) (*pb.Response, error) {
	consignment, err := s.repo.Create(req)
	if err != nil {
		fmt.Errorf("Cannot add consignment: %v", err)
		return nil, err
	}
	return &pb.Response{Created: true, Consignment: consignment}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Cannot listen port %v, get error %v", port, err)
	}
	s := grpc.NewServer()
	repo := &Repository{}
	pb.RegisterShippingServiceServer(s, &service{repo})
	reflection.Register(s)

	log.Println("Running on port:", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
