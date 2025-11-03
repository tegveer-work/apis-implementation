package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "grpc-api/proto"
)

type server struct {
	pb.UnimplementedEmployeeServiceServer
	employees []*pb.Employee
}

func (s *server) GetEmployees(ctx context.Context, req *pb.GetEmployeesRequest) (*pb.GetEmployeesResponse, error) {
	return &pb.GetEmployeesResponse{Employees: s.employees}, nil
}

func (s *server) AddEmployee(ctx context.Context, req *pb.AddEmployeeRequest) (*pb.AddEmployeeResponse, error) {
	newEmp := &pb.Employee{
		Id:      int32(len(s.employees) + 1),
		Name:    req.Name,
		Age:     req.Age,
		Company: req.Company,
		Health:  req.Health,
	}
	s.employees = append(s.employees, newEmp)
	return &pb.AddEmployeeResponse{Employee: newEmp}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterEmployeeServiceServer(grpcServer, &server{
		employees: []*pb.Employee{
			{Id: 1, Name: "Rohit", Age: 30, Company: "Lupin", Health: "Fit"},
			{Id: 2, Name: "Aarti", Age: 28, Company: "JK Tyres", Health: "Needs checkup"},
		},
	})

	fmt.Println("gRPC server running on port :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
