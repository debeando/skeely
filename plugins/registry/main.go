package registry

import (
	"mysql-ddl-lint/table"
)

type Method interface {
	Run(p Property)
}

type Creator func() Method

// List of handlers.
var Plugins = map[string]Creator{}

// Property has the particular properties for each handler.
type Property struct {
	FilePath string
	Table    table.Table
	Code     int
	Messages []Message
}

type Message struct {
	Code    int
	Message string
}

// Add new handler.
func Add(name string, creator Creator) {
	Plugins[name] = creator
}
