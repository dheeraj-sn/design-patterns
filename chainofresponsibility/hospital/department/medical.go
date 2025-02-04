package department

import (
	"fmt"
	"github.com/dheeraj-sn/design-patterns/chainofresponsibility/hospital/patient"
)

type Medical struct {
	Next Department
}

func (m *Medical) Execute(p *patient.Patient) {
	if p.MedicationDone {
		fmt.Println("Medical is already done")
		return
	}
	fmt.Println("Medical is providing medication")
	p.MedicationDone = true
	fmt.Println(p)
	if m.Next != nil {
		m.Next.Execute(p)
	} else {
		fmt.Println("Next department is not set for medical")
	}
}

func (m *Medical) SetNext(next Department) {
	m.Next = next
}
