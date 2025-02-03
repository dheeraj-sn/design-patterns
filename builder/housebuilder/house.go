package housebuilder

import "fmt"

type House struct {
	Type       string
	WindowType string
	DoorType   string
	Floor      int
}

func (h House) String() string {
	return fmt.Sprintf("%s: Window:%s, Door:%s, Floors:%d", h.Type, h.WindowType, h.DoorType, h.Floor)
}
