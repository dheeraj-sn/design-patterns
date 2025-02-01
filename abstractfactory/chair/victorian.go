package chair

import "fmt"

type Victorian struct {
}

func (v *Victorian) hasLegs() {
	fmt.Println("victorian chair hasLegs")
}

func (v *Victorian) sitOn() {
	fmt.Println("victorian chair sitOn")
}
