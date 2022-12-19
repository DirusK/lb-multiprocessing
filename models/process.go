package models

import (
	"fmt"
	"io"
	"os/exec"
)

type (
	Process struct {
		ID        uint
		Command   *exec.Cmd
		StdinPipe io.WriteCloser
	}

	Processes map[ProcessType][]Process
)

func NewProcess(processType ProcessType) (Process, error) {
	p := Process{
		ID: GetID(),
	}

	p.Command = exec.Command(
		"cmd.exe",
		"/C",
		"start",
		"D:\\Projects\\multiprocessing\\process\\process.exe",
		"-t",
		processType.String(),
	)

	var err error

	p.StdinPipe, err = p.Command.StdinPipe()
	if err != nil {
		return Process{}, err
	}

	if err = p.Command.Start(); err != nil {
		return Process{}, err
	}

	return p, nil
}

func (p Process) SendMessage(text string) error {
	if _, err := io.WriteString(p.StdinPipe, text); err != nil {
		fmt.Println("send message error: ", err)
		return err
	}

	return nil
}
