package coffeetable

import "fmt"

type ArtDeco struct {
}

func (v *ArtDeco) hasLegs() {
	fmt.Println("artDeco coffee table hasLegs")
}
