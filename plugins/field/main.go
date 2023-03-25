package field

import (
	"fmt"
	"strings"

	"skeely/config"
	"skeely/message"
	"skeely/registry"
)

var cnf *config.Config

type Plugin struct {
	Arguments registry.Arguments
	Messages  message.Messages
	Incidents message.Plugin
}

func init() {
	registry.Add(400, func() registry.Module { return &Plugin{} })

	cnf = config.GetInstance()
	cnf.Load()
}

func (p *Plugin) Run(a registry.Arguments) message.Plugin {
	p.Arguments = a
	p.Messages = []message.Message{
		{Code: 1, Message: "Table without fields."},
		{Code: 2, Message: "Table with many fields."},
		{Code: 3, Message: "Field name is to large, max 40: %s"},
		{Code: 4, Message: "Field name contains dot's, please remove it: %s"},
		{Code: 5, Message: "Field name has capital letter, please use lower case: %s"},
		{Code: 6, Message: "Field should by have comment: %s"},
		{Code: 7, Message: "Field with char type should by have length less than 50 chars: %s %s(%d)"},
		{Code: 8, Message: "Field varchar type with length great than 255 should by text type: %s %s(%d)"},
		{Code: 9, Message: "Field datetime type is defined, should by timestamp: %s"},
	}

	p.Empty()
	p.ManyFields()
	p.Length()
	p.Dots()
	p.LowerCase()
	p.Comment()
	p.CharLength()
	p.VarcharLength()
	p.HaveDatetime()

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
	p.Incidents.Add(message.Message{
		Code:    id,
		Message: fmt.Sprintf(p.GetMessage(id), vals...),
	})
}

func (p *Plugin) Empty() {
	if len(p.Arguments.Table.Fields) == 0 {
		p.AddMessage(1)
	}
}

func (p *Plugin) ManyFields() {
	if len(p.Arguments.Table.Fields) >= cnf.FieldsMax {
		p.AddMessage(2)
	}
}

func (p *Plugin) Length() {
	// https://dev.mysql.com/doc/refman/8.0/en/identifier-length.html
	for _, field := range p.Arguments.Table.Fields {
		if len(field.Name) >= 64 {
			p.AddMessage(3, field.Name)
		}
	}
}

func (p *Plugin) Dots() {
	for _, field := range p.Arguments.Table.Fields {
		if strings.Contains(field.Name, ".") {
			p.AddMessage(4, field.Name)
		}
	}
}

func (p *Plugin) LowerCase() {
	for _, field := range p.Arguments.Table.Fields {
		for _, r := range field.Name {
			if r >= 'A' && r <= 'Z' {
				p.AddMessage(5, field.Name)
			}
		}
	}
}

func (p *Plugin) Comment() {
	for _, field := range p.Arguments.Table.Fields {
		if len(field.Comment) == 0 {
			p.AddMessage(6, field.Name)
		}
	}
}

func (p *Plugin) CharLength() {
	for _, field := range p.Arguments.Table.Fields {
		if field.Type == "CHAR" && field.Length >= cnf.CharLengthMax {
			p.AddMessage(7, field.Name, field.Type, field.Length)
		}
	}
}

func (p *Plugin) VarcharLength() {
	for _, field := range p.Arguments.Table.Fields {
		if field.Type == "VARCHAR" && field.Length >= cnf.VarcharLengthMax {
			p.AddMessage(8, field.Name, field.Type, field.Length)
		}
	}
}

func (p *Plugin) HaveDatetime() {
	for _, field := range p.Arguments.Table.Fields {
		if field.Type == "DATETIME" {
			p.AddMessage(9, field.Name)
		}
	}
}
