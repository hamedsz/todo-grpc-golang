package main

import (
	"go-grpc/internal/models"
	"go-grpc/internal/rpc"
	"go-grpc/internal/services/todo"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	models.AutoMigrate()

	var port = ":9090"
	lis , err := net.Listen("tcp" , port)
	if err != nil{
		log.Fatalf("failed to listen: %v" , err)
	}

	s := grpc.NewServer()
	rpc.RegisterTodoServiceServer(s , &todo.TodoServiceServer{})

	log.Printf("server listening at %v" , lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve grpc: %v" , err)
	}
}
