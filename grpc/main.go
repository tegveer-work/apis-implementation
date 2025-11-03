package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "grpc-api/proto"
)

// server implements the SalaryService
type server struct {
	pb.UnimplementedSalaryServiceServer
}

// ComputeSalary calculates the total salary from base and bonus percent
func (s *server) ComputeSalary(ctx context.Context, req *pb.SalaryRequest) (*pb.SalaryResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("invalid request")
	}

	total := req.Base + (req.Base * req.BonusPercent / 100)
	msg := fmt.Sprintf("Salary computed for %s", req.Name)

	return &pb.SalaryResponse{
		Name:         req.Name,
		TotalSalary:  total,
		Message:      msg,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterSalaryServiceServer(s, &server{})

	fmt.Println("Salary gRPC service running on :50051 ...")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
