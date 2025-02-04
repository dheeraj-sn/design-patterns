package main

import (
	"fmt"
	"github.com/dheeraj-sn/design-patterns/abstractfactory/furniturefactory"
	"github.com/dheeraj-sn/design-patterns/builder/housebuilder"
	"github.com/dheeraj-sn/design-patterns/chainofresponsibility/hospital/department"
	"github.com/dheeraj-sn/design-patterns/chainofresponsibility/hospital/patient"
	"github.com/dheeraj-sn/design-patterns/prototype"
	"github.com/dheeraj-sn/design-patterns/singleton"
	"sync"
)

func main() {
	// abstractFactory()
	// buildHouseTest()
	// checkPrototype()
	//checkSingleton()
	patientGoesToHospital()

}

func patientGoesToHospital() {
	p := patient.Patient{Name: "John"}
	reception := &department.Reception{}
	medical := &department.Medical{}
	doctor := &department.Doctor{}
	cashier := &department.Cashier{}
	reception.SetNext(doctor)
	doctor.SetNext(medical)
	medical.SetNext(cashier)
	fmt.Println(p)
	reception.Execute(&p)
	fmt.Println(p)
}

func checkSingleton() {
	var wc sync.WaitGroup
	wc.Add(20)
	for i := 0; i < 20; i++ {
		go singleton.GetSingleObject(&wc, i)
	}
	wc.Wait()
	//fmt.Scanln()
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
