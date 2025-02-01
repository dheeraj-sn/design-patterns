package chair

import "fmt"

type ArtDeco struct {
}

func (v *ArtDeco) hasLegs() {
	fmt.Println("artDeco chair hasLegs")
}

func (v *ArtDeco) sitOn() {
	fmt.Println("artDeco chair sitOn")
}
