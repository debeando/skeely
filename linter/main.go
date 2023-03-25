package linter

import (
	"skeely/common"
	"skeely/config"
	"skeely/directory"
	"skeely/message"
	"skeely/registry"
	"skeely/table"
)

func Run() (allMessages message.Plugins) {
	cnf := config.GetInstance()

	directory.Iterator(func(fileName, fileContent string) {
		pluginResult := message.Plugin{File: fileName}
		tbl := table.Table{}

		if tbl.Parser(fileContent) != nil {
			return
		}

		// plugins inputs Iterator
		registry.Iterator(func(index int, creator registry.Creator) {
			// Create a plugin:
			plugin := creator()

			// Execute the plugin:
			messages := plugin.Run(
				registry.Arguments{
					Path:  fileName,
					Table: tbl,
				})

			// Process plugin output:
			messages.Iterator(func(msg message.Message) {
				if common.IntInSliceInt(cnf.IgnoreCodes(tbl.Name), index+msg.Code) {
					return
				}
				pluginResult.Add(message.Message{
					Code:    index + msg.Code,
					Message: msg.Message,
				})
			})
		})

		allMessages.Add(pluginResult)
	})

	return allMessages
}
