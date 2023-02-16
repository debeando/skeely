package field

import (
	"fmt"
	"strings"

	"mysql-ddl-lint/plugins/registry"
)

type Field struct {
	Property registry.Property
}

func init() {
	registry.Add("Field", func() registry.Method { return &Field{} })
}

func (f *Field) Run(p registry.Property) {
	f.Property = p
	f.Property.Code = 400

	f.Empty()
	f.ManyFields()
	f.Length()
	f.Dots()
	f.LowerCase()
	f.Comment()
	f.CharLength()
	f.VarcharLength()
	f.HaveDatetime()

	for _, message := range f.Property.Messages {
		fmt.Println(fmt.Sprintf("- [%d] %s", f.Property.Code + message.Code, message.Message))
	}
}

func (f *Field) AddMessage(id int, m string) {
	f.Property.Messages = append(f.Property.Messages, registry.Message{Code: id, Message: m})
}

func (f *Field) Empty() {
	if len(f.Property.Table.Fields) == 0 {
		f.AddMessage(1, "Table no have fields.")
	}
}

func (f *Field) ManyFields() {
	if len(f.Property.Table.Fields) >= 20 {
		f.AddMessage(2, "Table have many fields.")
	}
}

func (f *Field) Length() {
	for _, field := range f.Property.Table.Fields {
		if len(field.Name) >= 40 {
			f.AddMessage(3, fmt.Sprintf("Field name is to large, max 40: %s.", field.Name))
		}
	}
}

func (f *Field) Dots() {
	for _, field := range f.Property.Table.Fields {
		if strings.Contains(field.Name, ".") {
			f.AddMessage(4, fmt.Sprintf("Field name contains dot's, please remove: %s.", field.Name))
		}
	}
}

func (f *Field) LowerCase() {
	for _, field := range f.Property.Table.Fields {
		for _, r := range field.Name {
			if r >= 'A' && r <= 'Z' {
				f.AddMessage(5, fmt.Sprintf("Field name has capital letter, please use lower case: %s.", field.Name))
			}
		}
	}
}

func (f *Field) Comment() {
	for _, field := range f.Property.Table.Fields {
		if len(field.Comment) == 0 {
			f.AddMessage(6, fmt.Sprintf("Field should by have comment: %s.", field.Name))
		}
	}
}

func (f *Field) CharLength() {
	for _, field := range f.Property.Table.Fields {
		if field.Type == "char" && field.Length >= 50 {
			f.AddMessage(7, fmt.Sprintf("Field with char type should by have length less than 50 chars: %s(%d).", field.Type, field.Length))
		}
	}
}

func (f *Field) VarcharLength() {
	for _, field := range f.Property.Table.Fields {
		if field.Type == "varchar" && field.Length >= 255 {
			f.AddMessage(8, fmt.Sprintf("Field varchar type with length great than 255 should by text type: %s(%d).", field.Type, field.Length))
		}
	}
}

func (f *Field) HaveDatetime() {
	for _, field := range f.Property.Table.Fields {
		if field.Type == "datetime" {
			f.AddMessage(9, fmt.Sprintf("Field datetime type is defined, should by timestamp: %s.", field.Name))
		}
	}
}
