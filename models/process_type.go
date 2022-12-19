package models

type ProcessType string

const (
	ParentProcessType ProcessType = "parent"
	ChildProcessType  ProcessType = "child"
	UndefinedProcess  ProcessType = "undefined"
)

func (p ProcessType) String() string {
	return string(p)
}
