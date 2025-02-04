package department

import (
	"fmt"
	"github.com/dheeraj-sn/design-patterns/chainofresponsibility/hospital/patient"
)

type Cashier struct {
	Next Department
}

func (c *Cashier) Execute(p *patient.Patient) {
	if p.PaymentDone {
		fmt.Println("Payment is already done")
		return
	}
	fmt.Println("Patient is paying cashier")
	p.PaymentDone = true
	fmt.Println(p)
	if c.Next != nil {
		c.Next.Execute(p)
	} else {
		fmt.Println("Next department is not set for cashier")
	}
}

func (c *Cashier) SetNext(next Department) {
	c.Next = next
}
