package main

import (
	"fmt"
	"github.com/dheeraj-sn/design-patterns/abstractfactory/furniturefactory"
	"github.com/dheeraj-sn/design-patterns/builder/housebuilder"
	"github.com/dheeraj-sn/design-patterns/prototype"
)

func main() {
	// abstractFactory()
	// buildHouseTest()
	checkPrototype()

}

func checkPrototype() {
	file1 := prototype.File{
		Name: "file1",
	}
	file2 := prototype.File{
		Name: "file2",
	}
	file3 := prototype.File{
		Name: "file3",
	}

	folder1 := prototype.Folder{
		Name: "folder1",
		Children: []prototype.Inode{
			&file1,
		},
	}

	folder2 := prototype.Folder{
		Name: "folder2",
		Children: []prototype.Inode{
			&file2,
			&file3,
			&folder1,
		},
	}

	folder2.Show("    ")

	cloneFolder := folder2.Clone()
	cloneFolder.Show("    ")
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
