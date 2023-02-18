package field

import (
	"fmt"
	"strings"

	"mysql-ddl-lint/plugins/registry"
)

type Plugin struct {
	Property registry.Property
}

func init() {
	registry.Add("Field", func() registry.Method { return &Plugin{} })
}

func (p *Plugin) Run(a registry.Property) registry.Property {
	p.Property = a
	p.Property.Code = 400

	p.Empty()
	p.ManyFields()
	p.Length()
	p.Dots()
	p.LowerCase()
	p.Comment()
	p.CharLength()
	p.VarcharLength()
	p.HaveDatetime()

	return p.Property
}

func (p *Plugin) AddMessage(id int, m string) {
	// para listar todos los codigos y hacer una lista de lo que hace actualizada.
	// los mensajes deben estar en una variable.
	// p.Property.Checks[1] = "Field name is to large, max 40: %s"
	// Nosotros en el addMessage solo pasamos el 1 y los args.

	// aqui es donde valido si el mensaje entra o no dependiendo del fichero?
	p.Property.Messages = append(p.Property.Messages, registry.Message{Code: id, Message: m})
	// si se puede evitar este codigo en cada plugin seria genial.
}

func (p *Plugin) Empty() {
	if len(p.Property.Table.Fields) == 0 {
		p.AddMessage(1, "Table no have fields.")
	}
}

func (p *Plugin) ManyFields() {
	if len(p.Property.Table.Fields) >= 20 {
		p.AddMessage(2, "Table have many fields.")
	}
}

func (p *Plugin) Length() {
	for _, field := range p.Property.Table.Fields {
		if len(field.Name) >= 40 {
			p.AddMessage(3, fmt.Sprintf("Field name is to large, max 40: %s", field.Name))
		}
	}
}

func (p *Plugin) Dots() {
	for _, field := range p.Property.Table.Fields {
		if strings.Contains(field.Name, ".") {
			p.AddMessage(4, fmt.Sprintf("Field name contains dot's, please remove: %s", field.Name))
		}
	}
}

func (p *Plugin) LowerCase() {
	for _, field := range p.Property.Table.Fields {
		for _, r := range field.Name {
			if r >= 'A' && r <= 'Z' {
				p.AddMessage(5, fmt.Sprintf("Field name has capital letter, please use lower case: %s", field.Name))
			}
		}
	}
}

func (p *Plugin) Comment() {
	for _, field := range p.Property.Table.Fields {
		if len(field.Comment) == 0 {
			p.AddMessage(6, fmt.Sprintf("Field should by have comment: %s", field.Name))
		}
	}
}

func (p *Plugin) CharLength() {
	for _, field := range p.Property.Table.Fields {
		if field.Type == "char" && field.Length >= 50 {
			p.AddMessage(7, fmt.Sprintf("Field with char type should by have length less than 50 chars: %s(%d)", field.Type, field.Length))
		}
	}
}

func (p *Plugin) VarcharLength() {
	for _, field := range p.Property.Table.Fields {
		if field.Type == "varchar" && field.Length >= 255 {
			p.AddMessage(8, fmt.Sprintf("Field varchar type with length great than 255 should by text type: %s(%d)", field.Type, field.Length))
		}
	}
}

func (p *Plugin) HaveDatetime() {
	for _, field := range p.Property.Table.Fields {
		if field.Type == "datetime" {
			p.AddMessage(9, fmt.Sprintf("Field datetime type is defined, should by timestamp: %s", field.Name))
		}
	}
}
