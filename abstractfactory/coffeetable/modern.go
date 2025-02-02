package coffeetable

import "fmt"

type Modern struct {
}

func (v *Modern) HasLegs() {
	fmt.Println("modern coffee table hasLegs")
}

func (v *Modern) PlaceOn() {
	fmt.Println("modern coffee table placeOn")
}
