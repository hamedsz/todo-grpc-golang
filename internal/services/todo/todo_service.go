package todo

import (
	"context"
	"errors"
	"go-grpc/internal/db"
	"go-grpc/internal/models/todo"
	"go-grpc/internal/rpc"
	"time"
)

type TodoServiceServer struct {
	rpc.UnimplementedTodoServiceServer
	streams []rpc.TodoService_StreamIndexServer
}

func (s *TodoServiceServer) Index(_ context.Context, in *rpc.Empty) (*rpc.IndexResponse, error) {

	var response rpc.IndexResponse
	err := db.GetDB().Order("id desc").Find(&response.Items).Error

	return &response, err
}

func (s *TodoServiceServer) Create(_ context.Context, in *rpc.NewTodo) (*rpc.Todo, error) {
	if len(in.Title) == 0 {
		return nil , errors.New("title should not empty")
	}
	
	item := todo.Todo{
		Title:   in.Title,
		Content: in.Content,
	}

	err := db.GetDB().Create(
		&item,
	).Error

	if err != nil{
		return nil , err
	}

	err = s.SyncStreams()

	return &rpc.Todo{
		Id:      int32(item.ID),
		Title:   item.Title,
		Content: item.Content,
	}, err
}

func (s *TodoServiceServer) Update(_ context.Context, in *rpc.Todo) (*rpc.Todo, error) {
	err := db.GetDB().Save(in).Error
	return in, err
}

func (s *TodoServiceServer) Show(_ context.Context, in *rpc.TodoId) (*rpc.Todo, error) {
	var result rpc.Todo
	err := db.GetDB().First(&result , in.Id).Error

	return &result, err
}

func (s *TodoServiceServer) Delete(_ context.Context, in *rpc.TodoId) (*rpc.Empty, error) {
	var result rpc.Todo
	err := db.GetDB().First(&result , in.Id).Error
	if err != nil{
		return nil, err
	}

	err = db.GetDB().Delete(&result).Error

	if err != nil{
		return nil , err
	}

	err = s.SyncStreams()

	return &rpc.Empty{}, err
}

func (s *TodoServiceServer) StreamIndex(in *rpc.Empty , stream rpc.TodoService_StreamIndexServer) error {
	s.streams = append(s.streams , stream)
	err := s.SyncStreams()

	time.Sleep(10 * time.Second)
	return err
}

func (s *TodoServiceServer) SyncStreams() error {

	for _ , stream := range s.streams{
		var response rpc.IndexResponse
		err := db.GetDB().Order("id desc").Find(&response.Items).Error
		if err != nil{
			return err
		}
		stream.Send(&response)
	}

	return nil
}