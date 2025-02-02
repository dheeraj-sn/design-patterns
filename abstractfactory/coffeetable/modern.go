package coffeetable

import "fmt"

type Modern struct {
}

func (v *Modern) hasLegs() {
	fmt.Println("modern coffee table hasLegs")
}

func (v *Modern) placeOn() {
	fmt.Println("modern coffee table placeOn")
}
