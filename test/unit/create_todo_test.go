package unit

import (
	"context"
	"go-grpc/config"
	"go-grpc/internal/db"
	"go-grpc/internal/models"
	"go-grpc/internal/rpc"
	"go-grpc/internal/services/todo"
	"testing"
	"time"
)



func SetTestDatabase()  {
	config.Database.Name = "testing"
	db.RefreshDB()
	models.AutoMigrate()
	db.SetDB(db.GetDB().Begin())
}
func GetContext()  context.Context{
	ctx  , cancel := context.WithTimeout(context.Background() , time.Second)
	defer cancel()
	return ctx
}
func GetNewTodo() rpc.NewTodo {
	return rpc.NewTodo{
		Title: "test_title",
		Content: "test_content",
	}
}
func RollbackDB(t *testing.T)  {
	err := db.GetDB().Rollback().Error
	if err != nil{
		t.Errorf("rollback database failed: %v", err)
	}
}

func TestCreateWithOkData(t *testing.T)  {
	SetTestDatabase()

	newTodo := GetNewTodo()
	todoServer := todo.TodoServiceServer{}
	createdTodo , err := todoServer.Create(GetContext() , &newTodo)

	if err != nil{
		t.Errorf("insert new todo with todoServerServer Create method failed with error: %v" , err)
	}
	if createdTodo.Title != newTodo.Title {
		t.Error("create todo title is not match with requested todo")
	}
	if createdTodo.Content != newTodo.Content {
		t.Error("create todo content is not match with requested todo")
	}

	RollbackDB(t)
}

func TestCreateWithEmptyTitle(t *testing.T)  {
	SetTestDatabase()

	newTodo := rpc.NewTodo{
		Title: "",
	}
	todoServer := todo.TodoServiceServer{}
	_ , err := todoServer.Create(GetContext() , &newTodo)

	if err == nil{
		t.Error("when todo title is empty should not accept the request")

		RollbackDB(t)
	}
}

func TestCreateWithNoTitle(t *testing.T)  {
	SetTestDatabase()

	newTodo := rpc.NewTodo{}
	todoServer := todo.TodoServiceServer{}
	_ , err := todoServer.Create(GetContext() , &newTodo)

	if err == nil{
		t.Error("when todo title is not set should not accept the request")

		RollbackDB(t)
	}
}