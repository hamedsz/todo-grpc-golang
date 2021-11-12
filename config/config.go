package config

type DatabaseConfig struct {
	Name string
}

var Database = DatabaseConfig{
	Name: "todo",
}