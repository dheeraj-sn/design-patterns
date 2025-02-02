package chair

import "fmt"

type Victorian struct {
}

func (v *Victorian) HasLegs() {
	fmt.Println("victorian chair hasLegs")
}

func (v *Victorian) SitOn() {
	fmt.Println("victorian chair sitOn")
}
