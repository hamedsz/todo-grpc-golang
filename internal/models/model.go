package models

import (
	"go-grpc/internal/db"
	"go-grpc/internal/models/todo"
	"log"
)

func AutoMigrate()  {
	err := db.GetDB().AutoMigrate(&todo.Todo{})
	if err != nil{
		log.Fatalf("migration failed: %v" , err)
	}
}
