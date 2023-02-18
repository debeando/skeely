package primary_key

import (
	"fmt"

	"mysql-ddl-lint/common"
	"mysql-ddl-lint/plugins/registry"
)

type Plugin struct {
	Property registry.Property
}

func init() {
	registry.Add("PrimaryKey", func() registry.Method { return &Plugin{} })
}

func (p *Plugin) Run(a registry.Property) registry.Property {
	p.Property = a
	p.Property.Code = 500

	p.Empty()
	p.Name()
	p.NotNull()
	p.BigInt()
	p.Unsigned()
	p.AutoIncrement()

	return p.Property
}

func (p *Plugin) AddMessage(id int, m string) {
	p.Property.Messages = append(p.Property.Messages, registry.Message{Code: id, Message: m})
}

func (p *Plugin) Empty() {
	if len(p.Property.Table.PrimaryKey) == 0 {
		p.AddMessage(1, "Table no have Primary Key.")
	}
}

func (p *Plugin) Name() {
	if len(p.Property.Table.PrimaryKey) == 1 {
		for _, key := range p.Property.Table.PrimaryKey {
			if key != "id" {
				p.AddMessage(2, fmt.Sprintf("Primary Key field name should by id: %s", key))
				return
			}
		}
	}
}

func (p *Plugin) NotNull() {
	for _, key := range p.Property.Table.PrimaryKey {
		for _, field := range p.Property.Table.Fields {
			if key == field.Name && !field.NotNull {
				p.AddMessage(3, fmt.Sprintf("Primary Key field should by NOT NULL: %s", field.Name))
			}
		}
	}
}

func (p *Plugin) BigInt() {
	for _, key := range p.Property.Table.PrimaryKey {
		for _, field := range p.Property.Table.Fields {
			if key == field.Name && common.StringIn(field.Type, "INT") && field.Type != "BIGINT" {
				p.AddMessage(4, fmt.Sprintf("The primary key field must be BIGINT: %s %s", field.Name, field.Type))
			}
		}
	}
}

func (p *Plugin) Unsigned() {
	for _, key := range p.Property.Table.PrimaryKey {
		for _, field := range p.Property.Table.Fields {
			if key == field.Name && common.StringIn(field.Type, "INT") && !field.Unsigned {
				p.AddMessage(5, fmt.Sprintf("Primary Key field should by unsigned: %s", field.Name))
			}
		}
	}
}

func (p *Plugin) AutoIncrement() {
	for _, key := range p.Property.Table.PrimaryKey {
		for _, field := range p.Property.Table.Fields {
			if key == field.Name && common.StringIn(field.Type, "INT") && !field.AutoIncrement {
				p.AddMessage(6, fmt.Sprintf("Primary Key field should by auto increment: %s", field.Name))
			}
		}
	}
}
