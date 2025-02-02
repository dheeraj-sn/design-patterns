package furniturefactory

import (
	"github.com/dheeraj-sn/design-patterns/abstractfactory/chair"
	"github.com/dheeraj-sn/design-patterns/abstractfactory/coffeetable"
	"github.com/dheeraj-sn/design-patterns/abstractfactory/sofa"
)

type ArtDeco struct {
}

func NewArtDeco() FurnitureFactory {
	return &ArtDeco{}
}

func (v *ArtDeco) CreateChair() chair.Chair {
	return &chair.ArtDeco{}
}
func (v *ArtDeco) CreateCoffeeTable() coffeetable.CoffeeTable {
	return &coffeetable.ArtDeco{}
}
func (v *ArtDeco) CreateSofa() sofa.Sofa {
	return &sofa.ArtDeco{}
}
