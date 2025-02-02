package furniturefactory

import (
	"github.com/dheeraj-sn/design-patterns/abstractfactory/chair"
	"github.com/dheeraj-sn/design-patterns/abstractfactory/coffeetable"
	"github.com/dheeraj-sn/design-patterns/abstractfactory/sofa"
)

type Victorian struct {
}

func (v *Victorian) createChair() chair.Chair {
	return &chair.Victorian{}
}
func (v *Victorian) createCoffeeTable() coffeetable.CoffeeTable {
	return &coffeetable.Victorian{}
}
func (v *Victorian) createSofa() sofa.Sofa {
	return &sofa.Victorian{}
}
