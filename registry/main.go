package registry

import (
	"mysql-ddl-lint/table"
)

type Creator func() Module

var Plugins = make(map[int]Creator)

type Arguments struct {
	Path  string
	Table table.Table
}

type Message struct {
	Code    int
	Message string
}

type Module interface {
	Run(Arguments) []Message
}

func Add(id int, creator Creator) {
	Plugins[id] = creator
}
