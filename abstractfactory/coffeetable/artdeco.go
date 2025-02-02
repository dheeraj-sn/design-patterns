package coffeetable

import "fmt"

type ArtDeco struct {
}

func (v *ArtDeco) HasLegs() {
	fmt.Println("artDeco coffee table hasLegs")
}

func (v *ArtDeco) PlaceOn() {
	fmt.Println("artDeco coffee table placeOn")
}
