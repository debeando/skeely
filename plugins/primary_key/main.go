package primary_key

import (
	"fmt"
	"strings"

	"mysql-ddl-lint/plugins/registry"
)

type PrimaryKey struct {
	Property registry.Property
}

func init() {
	registry.Add("PrimaryKey", func() registry.Method { return &PrimaryKey{} })
}

func (pk *PrimaryKey) Run(p registry.Property) {
	pk.Property = p
	pk.Property.Code = 500

	pk.Empty()
	pk.Name()
	pk.NotNull()
	pk.BigInt()
	pk.Unsigned()
	pk.AutoIncrement()

	for _, message := range pk.Property.Messages {
		fmt.Println(fmt.Sprintf("- [%d] %s", pk.Property.Code+message.Code, message.Message))
	}
}

func (pk *PrimaryKey) AddMessage(id int, m string) {
	pk.Property.Messages = append(pk.Property.Messages, registry.Message{Code: id, Message: m})
}

func (pk *PrimaryKey) Empty() {
	if len(pk.Property.Table.PrimaryKey) == 0 {
		pk.AddMessage(1, "Table no have Primary Key.")
	}
}

func (pk *PrimaryKey) Name() {
	if len(pk.Property.Table.PrimaryKey) == 1 {
		for _, key := range pk.Property.Table.PrimaryKey {
			if key != "id" {
				pk.AddMessage(2, fmt.Sprintf("Primary Key field name should by id: %s", key))
				return
			}
		}
	}
}

func (pk *PrimaryKey) NotNull() {
	for _, key := range pk.Property.Table.PrimaryKey {
		for _, field := range pk.Property.Table.Fields {
			if key == field.Name && !field.NotNull {
				pk.AddMessage(3, fmt.Sprintf("Primary Key field should by NOT NULL: %s", field.Name))
			}
		}
	}
}

func (pk *PrimaryKey) BigInt() {
	for _, key := range pk.Property.Table.PrimaryKey {
		for _, field := range pk.Property.Table.Fields {
			if key == field.Name && strings.Contains(strings.ToUpper(field.Type), "INT") && strings.ToUpper(field.Type) != "BIGINT" {
				pk.AddMessage(4, fmt.Sprintf("Primary Key field should by BIGINT: %s %s", field.Name, field.Type))
			}
		}
	}
}

func (pk *PrimaryKey) Unsigned() {
	for _, key := range pk.Property.Table.PrimaryKey {
		for _, field := range pk.Property.Table.Fields {
			if key == field.Name && strings.Contains(strings.ToUpper(field.Type), "INT") && !field.Unsigned {
				pk.AddMessage(5, fmt.Sprintf("Primary Key field should by unsigned: %s", field.Name))
			}
		}
	}
}

func (pk *PrimaryKey) AutoIncrement() {
	for _, key := range pk.Property.Table.PrimaryKey {
		for _, field := range pk.Property.Table.Fields {
			if key == field.Name && strings.Contains(strings.ToUpper(field.Type), "INT") && !field.AutoIncrement {
				pk.AddMessage(6, fmt.Sprintf("Primary Key field should by auto increment: %s", field.Name))
			}
		}
	}
}
