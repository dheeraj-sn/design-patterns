package department

import "github.com/dheeraj-sn/design-patterns/chainofresponsibility/hospital/patient"

type Department interface {
	Execute(*patient.Patient)
	SetNext(Department)
}
