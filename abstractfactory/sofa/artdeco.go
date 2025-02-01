package sofa

import "fmt"

type ArtDeco struct {
}

func (v *ArtDeco) hasLegs() {
	fmt.Println("artDeco sofa hasLegs")
}
