package todo

import (
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Title string
	Content string
}

