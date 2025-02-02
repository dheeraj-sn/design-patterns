package furniturefactory

import (
	"github.com/dheeraj-sn/design-patterns/abstractfactory/chair"
	"github.com/dheeraj-sn/design-patterns/abstractfactory/coffeetable"
	"github.com/dheeraj-sn/design-patterns/abstractfactory/sofa"
)

type ArtDeco struct {
}

func (v *ArtDeco) createChair() chair.Chair {
	return &chair.ArtDeco{}
}
func (v *ArtDeco) createCoffeeTable() coffeetable.CoffeeTable {
	return &coffeetable.ArtDeco{}
}
func (v *ArtDeco) createSofa() sofa.Sofa {
	return &sofa.ArtDeco{}
}
