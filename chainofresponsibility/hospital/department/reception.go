package department

import (
	"fmt"
	"github.com/dheeraj-sn/design-patterns/chainofresponsibility/hospital/patient"
)

type Reception struct {
	Next Department
}

func (r *Reception) Execute(p *patient.Patient) {
	if p.RegistrationDone {
		fmt.Println("Patient registration is already done")
		return
	}
	fmt.Println("Reception is registering patient")
	p.RegistrationDone = true
	fmt.Println(p)
	if r.Next != nil {
		r.Next.Execute(p)
	} else {
		fmt.Println("Next department for reception is not set")
	}
}

func (r *Reception) SetNext(d Department) {
	r.Next = d
}
