package chair

import "fmt"

type Modern struct {
}

func (v *Modern) hasLegs() {
	fmt.Println("modern chair hasLegs")
}

func (v *Modern) sitOn() {
	fmt.Println("modern chair sitOn")
}
