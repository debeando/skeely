package primary_key

import (
	"fmt"

	"mylinter/common"
	"mylinter/registry"
)

type Plugin struct {
	Arguments registry.Arguments
	Messages  []registry.Message
	Incidents []registry.Message
}

func init() {
	registry.Add(500, func() registry.Module { return &Plugin{} })
}

func (p *Plugin) Run(a registry.Arguments) []registry.Message {
	p.Arguments = a
	p.Messages = []registry.Message{
		{Code: 1, Message: "Table no have Primary Key."},
		{Code: 2, Message: "Primary Key field name should by id: %s"},
		{Code: 3, Message: "Primary Key field should by NOT NULL: %s"},
		{Code: 4, Message: "Primary key field must be BIGINT: %s %s"},
		{Code: 5, Message: "Primary Key field should by unsigned: %s"},
		{Code: 6, Message: "Primary Key field should by auto increment: %s"},
	}

	p.Empty()
	p.Name()
	p.NotNull()
	p.BigInt()
	p.Unsigned()
	p.AutoIncrement()

	return p.Incidents
}

func (p *Plugin) GetMessage(id int) string {
	for _, message := range p.Messages {
		if message.Code == id {
			return message.Message
		}
	}
	return ""
}

func (p *Plugin) AddMessage(id int, vals ...any) {
	msg := registry.Message{Code: id}
	msg.Message = fmt.Sprintf(p.GetMessage(id), vals...)

	p.Incidents = append(p.Incidents, msg)
}

func (p *Plugin) Empty() {
	if len(p.Arguments.Table.PrimaryKey) == 0 {
		p.AddMessage(1, "")
	}
}

func (p *Plugin) Name() {
	if len(p.Arguments.Table.PrimaryKey) == 1 {
		for _, key := range p.Arguments.Table.PrimaryKey {
			if key != "id" {
				p.AddMessage(2, key)
				return
			}
		}
	}
}

func (p *Plugin) NotNull() {
	for _, key := range p.Arguments.Table.PrimaryKey {
		for _, field := range p.Arguments.Table.Fields {
			if key == field.Name && !field.NotNull {
				p.AddMessage(3, field.Name)
			}
		}
	}
}

func (p *Plugin) BigInt() {
	for _, key := range p.Arguments.Table.PrimaryKey {
		for _, field := range p.Arguments.Table.Fields {
			if key == field.Name && common.StringIn(field.Type, "INT") && field.Type != "BIGINT" {
				p.AddMessage(4, field.Name, field.Type)
			}
		}
	}
}

func (p *Plugin) Unsigned() {
	for _, key := range p.Arguments.Table.PrimaryKey {
		for _, field := range p.Arguments.Table.Fields {
			if key == field.Name && common.StringIn(field.Type, "INT") && !field.Unsigned {
				p.AddMessage(5, field.Name)
			}
		}
	}
}

func (p *Plugin) AutoIncrement() {
	for _, key := range p.Arguments.Table.PrimaryKey {
		for _, field := range p.Arguments.Table.Fields {
			if key == field.Name && common.StringIn(field.Type, "INT") && !field.AutoIncrement {
				p.AddMessage(6, field.Name)
			}
		}
	}
}
