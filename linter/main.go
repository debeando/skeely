package linter

import (
	"skeely/common"
	"skeely/config"
	"skeely/registry"
	"skeely/table"
	"skeely/directory"
)

type Linter struct {
	Path     string
	Summary  []Result
}

type Result struct {
	File string
	Messages []Message
}

type Message struct {
	Code int
	Message string
}

var instance *Linter

func GetInstance() *Linter {
	if instance == nil {
		instance = &Linter{}
	}
	return instance
}

func (l *Linter) Run() {
	cnf := config.GetInstance()
	cnf.Load()

	directory.Explore(l.Path, func(fileName, fileContent string) {
		r := Result{File: fileName}
		t := table.Table{}

		if t.Parser(fileContent) != nil {
			return
		}

		for key := range registry.Plugins {
			if creator, ok := registry.Plugins[key]; ok {
				plugin := creator()
				messages := plugin.Run(registry.Arguments{Path: fileName, Table: t})

				for _, message := range messages {
					if common.IntInArrayInt(cnf.IgnoreCodes(t.Name), key+message.Code) {
						continue
					}

					r.AddMessage(Message{Code: key+message.Code, Message: message.Message})
				}
			}
		}
		l.AddResult(r)
	})
}

func (l *Linter) AddResult(r Result) {
	l.Summary = append(l.Summary, r)
}


func (r *Result) AddMessage(m Message) {
	r.Messages = append(r.Messages, m)
}