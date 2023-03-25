package registry

import (
	"skeely/message"
	"skeely/table"
)

type Creator func() Module

var Plugins = make(map[int]Creator)

type Arguments struct {
	Path  string
	Table table.Table
}

type Module interface {
	Run(Arguments) message.Plugin
}

func Add(id int, creator Creator) {
	Plugins[id] = creator
}

func Iterator(doIterator func(int, Creator)) {
	for index := range Plugins {
		if creator, ok := Plugins[index]; ok {
			doIterator(index, creator)
		}
	}
}
