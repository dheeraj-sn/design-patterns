package coffeetable

import "fmt"

type Victorian struct {
}

func (v *Victorian) hasLegs() {
	fmt.Println("victorian coffee table hasLegs")
}

func (v *Victorian) placeOn() {
	fmt.Println("victorian coffee table placeOn")
}
