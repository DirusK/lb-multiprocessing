package boss

import (
	"fmt"
	"time"

	"multiprocessing/models"
)

type (
	Boss struct {
		processes       models.Processes
		messagesChannel chan []Message
		capacity        MessageCapacity
	}

	Message struct {
		ProcessType models.ProcessType
		Message     string
	}
)

func New() *Boss {
	boss := Boss{
		processes:       make(models.Processes),
		messagesChannel: make(chan []Message),
		capacity:        NewMessageCapacity(getCapacity()),
	}

	var (
		childProcesses  = getNumberOfChildProcesses()
		parentProcesses = getNumberOfParentProcesses()
	)

	for i := 0; i < childProcesses; i++ {
		process, err := models.NewProcess(models.ChildProcessType)
		if err != nil {
			panic(err)
		}

		boss.processes[models.ChildProcessType] = append(boss.processes[models.ChildProcessType], process)
	}

	for i := 0; i < parentProcesses; i++ {
		process, err := models.NewProcess(models.ParentProcessType)
		if err != nil {
			panic(err)
		}

		boss.processes[models.ParentProcessType] = append(boss.processes[models.ParentProcessType], process)
	}

	return &boss
}

// Run starts the boss.
func (b *Boss) Run() {
	go b.listenToMessages()
	for {
		b.executeByMenu()
		time.Sleep(time.Second * 10)
	}
}

func (b *Boss) listenToMessages() {
	var (
		childsLen  = len(b.processes[models.ChildProcessType])
		parentsLen = len(b.processes[models.ParentProcessType])
	)

	var (
		childCounter  = 0
		parentCounter = 0
	)

	for {
		messages := <-b.messagesChannel

		for _, message := range messages {
			message := message

			go func() {
				if b.capacity.IsOverflowed() {
					fmt.Println("Capacity is overflowed, drops the message")
					return
				}

				switch message.ProcessType {
				case models.ChildProcessType:
					b.processes[models.ChildProcessType][childCounter].SendMessage(message.Message)

					childCounter++
					if childCounter == childsLen {
						childCounter = 0
					}

				case models.ParentProcessType:
					b.processes[models.ParentProcessType][parentCounter].SendMessage(message.Message)

					parentCounter++
					if parentCounter == parentsLen {
						parentCounter = 0
					}
				}
			}()
		}

		b.capacity.Refresh()
	}
}
