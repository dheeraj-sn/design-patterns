package chair

import "fmt"

type Modern struct {
}

func (v *Modern) HasLegs() {
	fmt.Println("modern chair hasLegs")
}

func (v *Modern) SitOn() {
	fmt.Println("modern chair sitOn")
}
