package sofa

import "fmt"

type Victorian struct {
}

func (v *Victorian) HasLegs() {
	fmt.Println("victorian sofa hasLegs")
}
