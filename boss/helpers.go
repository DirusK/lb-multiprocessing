package boss

import (
	"errors"
	"fmt"
	"os"

	"multiprocessing/models"
)

const defaultCapacity = 3

func (b *Boss) executeByMenu() {
	switch b.menuChoice() {
	case models.SendMessageToChilds:
		var messages []Message
		for i := 0; i < len(b.processes[models.ChildProcessType]); i++ {
			messages = append(messages, Message{
				ProcessType: models.ChildProcessType,
				Message:     "Hello child",
			})
		}

		fmt.Printf("Sending %d messages to childs... \n", len(messages))
		b.messagesChannel <- messages

	case models.SendMessageToParents:
		var messages []Message
		for i := 0; i < len(b.processes[models.ParentProcessType]); i++ {
			messages = append(messages, Message{
				ProcessType: models.ParentProcessType,
				Message:     "Hello parent",
			})
		}

		fmt.Printf("Sending %d messages to parents... \n", len(messages))
		b.messagesChannel <- messages

	case models.SendMessageToChildsAndParents:
		var messages []Message

		for i := 0; i < len(b.processes[models.ChildProcessType]); i++ {
			messages = append(messages, Message{
				ProcessType: models.ChildProcessType,
				Message:     "Hello child",
			})
		}

		for i := 0; i < len(b.processes[models.ParentProcessType]); i++ {
			messages = append(messages, Message{
				ProcessType: models.ParentProcessType,
				Message:     "Hello parent",
			})
		}

		fmt.Printf("Sending %d messages to childs and parents... \n", len(messages))
		b.messagesChannel <- messages

	case models.Exit:
		var messages []Message

		for i := 0; i < len(b.processes[models.ChildProcessType]); i++ {
			messages = append(messages, Message{
				ProcessType: models.ChildProcessType,
				Message:     models.MessageExit,
			})
		}

		for i := 0; i < len(b.processes[models.ParentProcessType]); i++ {
			messages = append(messages, Message{
				ProcessType: models.ParentProcessType,
				Message:     models.MessageExit,
			})
		}

		fmt.Printf("Sending %d messages to childs and parents... \n", len(messages))
		b.messagesChannel <- messages

		fmt.Println("Exiting...")
		fmt.Println("Sending message to childs and parents to exit...")

		os.Exit(0)
	}
}

// menuChoice opens the menu.
func (b *Boss) menuChoice() int {
	fmt.Println("Choose an option:")
	fmt.Printf("%d - Send message to childs \n", models.SendMessageToChilds)
	fmt.Printf("%d - Send message to parents \n", models.SendMessageToParents)
	fmt.Printf("%d - Send message to childs and parents \n", models.SendMessageToChildsAndParents)
	fmt.Printf("%d - Exit \n", models.Exit)

	var choice int

	for {
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		if choice >= models.SendMessageToChilds && choice <= models.Exit {
			break
		}

		fmt.Println("ERROR: ", errors.New("invalid choice. Try again"))
	}

	return choice
}

// getCapacity returns the capacity of the message buffer.
func getCapacity() int {
	var capacity int

	fmt.Print("Enter the capacity of messages: ")
	fmt.Scanln(&capacity)

	if capacity == 0 {
		capacity = defaultCapacity
	}

	return capacity
}

// getNumberOfParentProcesses returns the number of parent processes.
func getNumberOfParentProcesses() int {
	var numberOfProcesses int

	fmt.Print("Enter the number of parent processes: ")
	fmt.Scanln(&numberOfProcesses)

	if numberOfProcesses == 0 {
		fmt.Println("ERROR: ", "The number of parent processes must be greater than 0.")
	}

	return numberOfProcesses
}

// getNumberOfChildProcesses returns the number of child processes.
func getNumberOfChildProcesses() int {
	var numberOfProcesses int

	fmt.Print("Enter the number of child processes: ")
	fmt.Scanln(&numberOfProcesses)

	if numberOfProcesses == 0 {
		fmt.Println("ERROR: ", "The number of child processes must be greater than 0.")
	}

	return numberOfProcesses
}
