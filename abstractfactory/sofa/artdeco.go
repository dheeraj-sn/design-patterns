package sofa

import "fmt"

type ArtDeco struct {
}

func (v *ArtDeco) HasLegs() {
	fmt.Println("artDeco sofa hasLegs")
}
