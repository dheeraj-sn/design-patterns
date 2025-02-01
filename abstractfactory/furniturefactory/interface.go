package furniturefactory

import (
	"github.com/dheeraj-sn/design-patterns/abstractfactory/chair"
	"github.com/dheeraj-sn/design-patterns/abstractfactory/coffeetable"
	"github.com/dheeraj-sn/design-patterns/abstractfactory/sofa"
)

type FurnitureFactory interface {
	createChair() chair.Chair
	createCoffeeTable() coffeetable.CoffeeTable
	createSofa() sofa.Sofa
}
