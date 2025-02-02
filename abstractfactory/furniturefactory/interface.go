package furniturefactory

import (
	"github.com/dheeraj-sn/design-patterns/abstractfactory/chair"
	"github.com/dheeraj-sn/design-patterns/abstractfactory/coffeetable"
	"github.com/dheeraj-sn/design-patterns/abstractfactory/sofa"
)

type FurnitureFactory interface {
	CreateChair() chair.Chair
	CreateCoffeeTable() coffeetable.CoffeeTable
	CreateSofa() sofa.Sofa
}
