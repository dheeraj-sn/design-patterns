package furniturefactory

import (
	"github.com/dheeraj-sn/design-patterns/abstractfactory/chair"
	"github.com/dheeraj-sn/design-patterns/abstractfactory/coffeetable"
	"github.com/dheeraj-sn/design-patterns/abstractfactory/sofa"
)

type Victorian struct {
}

func NewVictorian() FurnitureFactory {
	return &Victorian{}
}

func (v *Victorian) CreateChair() chair.Chair {
	return &chair.Victorian{}
}
func (v *Victorian) CreateCoffeeTable() coffeetable.CoffeeTable {
	return &coffeetable.Victorian{}
}
func (v *Victorian) CreateSofa() sofa.Sofa {
	return &sofa.Victorian{}
}
