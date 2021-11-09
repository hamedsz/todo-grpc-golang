package main

import (
	"context"
	"go-grpc/internal/rpc"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	address = "localhost:6565"
)

func main() {
	conn , err := grpc.Dial(address , grpc.WithInsecure() , grpc.WithBlock())
	if err != nil{
		log.Fatalf("did not connect: %v" , err)
	}
	defer conn.Close()

	c := rpc.NewTodoServiceClient(conn)
	ctx  , cancel := context.WithTimeout(context.Background() , time.Second)
	defer cancel()

	createTodo(c , ctx)
	indexTodos(c , ctx)
}

func createTodo(c rpc.TodoServiceClient , ctx context.Context)  {

	result , err := c.Create(ctx , &rpc.NewTodo{
		Title: "hello world",
		Content: "hello world content",
	})
	if err != nil {
		log.Fatalf("could not create todo: %v" , err)
	}
	log.Printf("todo detail: %v" , result)
}

func indexTodos(c rpc.TodoServiceClient , ctx context.Context)  {

	result , err := c.Index(ctx , &rpc.Empty{})
	if err != nil {
		log.Fatalf("could not index todo")
	}

	log.Println("list todos")
	log.Println(result)
}