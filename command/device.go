package command

import "fmt"

type Device interface {
	On()
	Off()
}

type TV struct {
	IsRunning bool
}

func (t *TV) On() {
	t.IsRunning = true
	fmt.Println("TV On")
}

func (t *TV) Off() {
	t.IsRunning = false
	fmt.Println("TV Off")
}
