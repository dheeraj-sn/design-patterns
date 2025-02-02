package coffeetable

import "fmt"

type ArtDeco struct {
}

func (v *ArtDeco) hasLegs() {
	fmt.Println("artDeco coffee table hasLegs")
}

func (v *ArtDeco) placeOn() {
	fmt.Println("artDeco coffee table placeOn")
}
