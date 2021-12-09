package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	pb "github.com/DymaSV/shippy-service-consignment/proto/consignment"
	"google.golang.org/grpc"
)

const (
	address  = "localhost:50051"
	filename = "consignment.json"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewShippingServiceClient(conn)

	file := filename
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	consignment, err := parseFile(file)
	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}

	r, err := client.CreateConsignment(context.Background(), consignment)
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Printf("Created: %t", r.Created)
}

func parseFile(filename string) (*pb.Consignment, error) {
	var consignment *pb.Consignment
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Cannot read file %v", err)
		return nil, err
	}
	err = json.Unmarshal(data, &consignment)
	if err != nil {
		fmt.Printf("Cannot read json %v", err)
		return nil, err
	}
	return consignment, err
}
