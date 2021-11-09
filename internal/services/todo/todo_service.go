package todo

import (
	"context"
	"go-grpc/internal/db"
	"go-grpc/internal/models/todo"
	"go-grpc/internal/rpc"
)

type TodoServiceServer struct {
	rpc.UnimplementedTodoServiceServer
}

func (s *TodoServiceServer) Index(ctx context.Context , in *rpc.Empty) (*rpc.IndexResponse , error)  {

	var response rpc.IndexResponse
	db.GetDB().Find(&response.Items)

	return &response , nil
}

func (s *TodoServiceServer) Create(ctx context.Context , in *rpc.NewTodo) (*rpc.Todo , error)  {
	db.GetDB().AutoMigrate(&todo.Todo{})

	todo := todo.Todo{
		Title: in.Title,
		Content: in.Content,
	}

	db.GetDB().Create(
		&todo,
	)

	return &rpc.Todo{
		Id: int32(todo.ID),
		Title: todo.Title,
		Content: todo.Content,
	} , nil
}

func (s *TodoServiceServer) Update(ctx context.Context , in *rpc.Todo) (*rpc.Todo , error) {
	return nil , nil
}

func (s *TodoServiceServer) Show(ctx context.Context , in *rpc.TodoId) (*rpc.Todo , error) {
	return nil , nil
}

func (s *TodoServiceServer) Delete(ctx context.Context , in *rpc.TodoId) (*rpc.Empty , error) {
	return nil , nil
}