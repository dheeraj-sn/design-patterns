package coffeetable

import "fmt"

type Victorian struct {
}

func (v *Victorian) HasLegs() {
	fmt.Println("victorian coffee table hasLegs")
}

func (v *Victorian) PlaceOn() {
	fmt.Println("victorian coffee table placeOn")
}
