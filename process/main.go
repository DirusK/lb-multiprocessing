package main

import (
	"fmt"
	"strings"

	"github.com/spf13/pflag"

	"multiprocessing/models"
)

func main() {
	processType := pflag.StringP("type", "t", models.UndefinedProcess.String(), "process type")
	pflag.Parse()

	fmt.Println("Process type: ", strings.ToUpper(*processType))
	fmt.Println("Waiting for messages...")

	for {
		var message string
		fmt.Scanln(&message)

		if message == models.MessageExit {
			fmt.Println("Exiting...")
			return
		}

		fmt.Println("Message received: ", message)
	}
}
