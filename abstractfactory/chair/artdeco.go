package chair

import "fmt"

type ArtDeco struct {
}

func (v *ArtDeco) HasLegs() {
	fmt.Println("artDeco chair hasLegs")
}

func (v *ArtDeco) SitOn() {
	fmt.Println("artDeco chair sitOn")
}
