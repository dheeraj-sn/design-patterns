package main

import (
	"fmt"
	"github.com/dheeraj-sn/design-patterns/abstractfactory/furniturefactory"
	"github.com/dheeraj-sn/design-patterns/builder/housebuilder"
)

func main() {
	// abstractFactory()
	buildHouseTest()

}

func buildHouseTest() {
	normalBuilder := housebuilder.NewNormalBuilder()
	iglooBuilder := housebuilder.NewIglooBuilder()
	director := housebuilder.NewDirector(normalBuilder)
	normalHouse := director.BuildHouse()
	fmt.Println(normalHouse)

	director.SetBuilder(iglooBuilder)
	iglooHouse := director.BuildHouse()
	fmt.Println(iglooHouse)

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
