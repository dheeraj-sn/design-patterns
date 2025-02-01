package coffeetable

import "fmt"

type Modern struct {
}

func (v *Modern) hasLegs() {
	fmt.Println("modern coffee table hasLegs")
}
