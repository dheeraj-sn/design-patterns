package sofa

import "fmt"

type Modern struct {
}

func (v *Modern) HasLegs() {
	fmt.Println("modern sofa hasLegs")
}
