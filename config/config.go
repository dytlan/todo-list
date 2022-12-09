package config

import (
	"github.com/dytlan/moonlay-todo-list/engine/file"
	"github.com/dytlan/moonlay-todo-list/engine/sql"
)

type Map struct {
	Database sql.Config
	Server   ServerConfig
	File     file.Config
}

type ServerConfig struct {
	Port string
}
