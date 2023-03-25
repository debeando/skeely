package terminal

import (
	"fmt"

	"skeely/message"
)

func Print(messages message.Plugins) {
	for _, msgPlugin := range messages {
		fmt.Println(fmt.Sprintf("> File: %s", msgPlugin.File))
		for _, msg := range msgPlugin.Messages {
			fmt.Println(fmt.Sprintf("- [%d] %s", msg.Code, msg.Message))
		}
		if len(msgPlugin.Messages) == 0 {
			fmt.Println("- Looks ok.")
		}
	}
}
