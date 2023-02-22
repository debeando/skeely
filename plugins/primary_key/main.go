package primary_key

import (
	"fmt"

	"mylinter/common"
	"mylinter/registry"
)

type Plugin struct {
	Arguments registry.Arguments
	Messages  []registry.Message
}

func init() {
	registry.Add(500, func() registry.Module { return &Plugin{} })
}

func (p *Plugin) Run(a registry.Arguments) []registry.Message {
	p.Arguments = a

	p.Empty()
	p.Name()
	p.NotNull()
	p.BigInt()
	p.Unsigned()
	p.AutoIncrement()

	return p.Messages
}

func (p *Plugin) AddMessage(id int, m string) {
	p.Messages = append(p.Messages, registry.Message{Code: id, Message: m})
}

func (p *Plugin) Empty() {
	if len(p.Arguments.Table.PrimaryKey) == 0 {
		p.AddMessage(1, "Table no have Primary Key.")
	}
}

func (p *Plugin) Name() {
	if len(p.Arguments.Table.PrimaryKey) == 1 {
		for _, key := range p.Arguments.Table.PrimaryKey {
			if key != "id" {
				p.AddMessage(2, fmt.Sprintf("Primary Key field name should by id: %s", key))
				return
			}
		}
	}
}

func (p *Plugin) NotNull() {
	for _, key := range p.Arguments.Table.PrimaryKey {
		for _, field := range p.Arguments.Table.Fields {
			if key == field.Name && !field.NotNull {
				p.AddMessage(3, fmt.Sprintf("Primary Key field should by NOT NULL: %s", field.Name))
			}
		}
	}
}

func (p *Plugin) BigInt() {
	for _, key := range p.Arguments.Table.PrimaryKey {
		for _, field := range p.Arguments.Table.Fields {
			if key == field.Name && common.StringIn(field.Type, "INT") && field.Type != "BIGINT" {
				p.AddMessage(4, fmt.Sprintf("Primary key field must be BIGINT: %s %s", field.Name, field.Type))
			}
		}
	}
}

func (p *Plugin) Unsigned() {
	for _, key := range p.Arguments.Table.PrimaryKey {
		for _, field := range p.Arguments.Table.Fields {
			if key == field.Name && common.StringIn(field.Type, "INT") && !field.Unsigned {
				p.AddMessage(5, fmt.Sprintf("Primary Key field should by unsigned: %s", field.Name))
			}
		}
	}
}

func (p *Plugin) AutoIncrement() {
	for _, key := range p.Arguments.Table.PrimaryKey {
		for _, field := range p.Arguments.Table.Fields {
			if key == field.Name && common.StringIn(field.Type, "INT") && !field.AutoIncrement {
				p.AddMessage(6, fmt.Sprintf("Primary Key field should by auto increment: %s", field.Name))
			}
		}
	}
}
