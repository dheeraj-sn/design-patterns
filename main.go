package main

import (
	"github.com/dheeraj-sn/design-patterns/abstractfactory/furniturefactory"
)

func main() {
	mf := furniturefactory.NewModern()
	mfc := mf.CreateChair()
	mfs := mf.CreateSofa()
	mfct := mf.CreateCoffeeTable()
	mfc.HasLegs()
	mfs.HasLegs()
	mfct.HasLegs()
}
