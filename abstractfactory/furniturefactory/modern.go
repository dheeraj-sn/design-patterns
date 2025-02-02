package furniturefactory

import (
	"github.com/dheeraj-sn/design-patterns/abstractfactory/chair"
	"github.com/dheeraj-sn/design-patterns/abstractfactory/coffeetable"
	"github.com/dheeraj-sn/design-patterns/abstractfactory/sofa"
)

type Modern struct {
}

func (v *Modern) createChair() chair.Chair {
	return &chair.Modern{}
}
func (v *Modern) createCoffeeTable() coffeetable.CoffeeTable {
	return &coffeetable.Modern{}
}
func (v *Modern) createSofa() sofa.Sofa {
	return &sofa.Modern{}
}
