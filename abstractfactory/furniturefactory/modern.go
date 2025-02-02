package furniturefactory

import (
	"github.com/dheeraj-sn/design-patterns/abstractfactory/chair"
	"github.com/dheeraj-sn/design-patterns/abstractfactory/coffeetable"
	"github.com/dheeraj-sn/design-patterns/abstractfactory/sofa"
)

type Modern struct {
}

func NewModern() FurnitureFactory {
	return &Modern{}
}

func (v *Modern) CreateChair() chair.Chair {
	return &chair.Modern{}
}
func (v *Modern) CreateCoffeeTable() coffeetable.CoffeeTable {
	return &coffeetable.Modern{}
}
func (v *Modern) CreateSofa() sofa.Sofa {
	return &sofa.Modern{}
}
