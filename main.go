package main

import (
	"fmt"
	"github.com/dheeraj-sn/design-patterns/abstractfactory/furniturefactory"
)

func main() {
	abstractFactory()

}

func abstractFactory() {
	mf := furniturefactory.NewModern()
	mfc := mf.CreateChair()
	mfs := mf.CreateSofa()
	mfct := mf.CreateCoffeeTable()
	mfc.HasLegs()
	mfc.SitOn()
	mfs.HasLegs()
	mfct.HasLegs()
	mfct.PlaceOn()

	fmt.Println("============")
	mf = furniturefactory.NewArtDeco()
	mfc = mf.CreateChair()
	mfs = mf.CreateSofa()
	mfct = mf.CreateCoffeeTable()
	mfc.HasLegs()
	mfc.SitOn()
	mfs.HasLegs()
	mfct.HasLegs()
	mfct.PlaceOn()
	fmt.Println("============")
	mf = furniturefactory.NewVictorian()
	mfc = mf.CreateChair()
	mfs = mf.CreateSofa()
	mfct = mf.CreateCoffeeTable()
	mfc.HasLegs()
	mfc.SitOn()
	mfs.HasLegs()
	mfct.HasLegs()
	mfct.PlaceOn()
}
