package main

import (
	"context"
	"fmt"
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

	before := time.Now().UnixNano() / int64(time.Millisecond)
	showTodo(c , ctx)
	after := time.Now().UnixNano() / int64(time.Millisecond)

	fmt.Println(after - before)
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

func deleteTodo(c rpc.TodoServiceClient , ctx context.Context)  {
	_ , err := c.Delete(ctx , &rpc.TodoId{
		Id: 1,
	})
	if err != nil {
		log.Fatalf("could not delete todo")
	}
}

func showTodo(c rpc.TodoServiceClient , ctx context.Context)  {
	result , err := c.Show(ctx , &rpc.TodoId{
		Id: 2,
	})
	if err != nil {
		log.Fatalf("could not delete todo")
	}

	fmt.Println(result)
}