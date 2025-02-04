package department

import (
	"fmt"
	"github.com/dheeraj-sn/design-patterns/chainofresponsibility/hospital/patient"
)

type Doctor struct {
	Next Department
}

func (d *Doctor) Execute(p *patient.Patient) {
	if p.DoctorCheckupDone {
		fmt.Println("Doctor checkup is already done")
		return
	}
	fmt.Println("Doctor is checking patient")
	p.DoctorCheckupDone = true
	fmt.Println(p)
	if d.Next != nil {
		d.Next.Execute(p)
	} else {
		fmt.Println("Next department for doctor is not set")
	}
}

func (d *Doctor) SetNext(next Department) {
	d.Next = next
}
