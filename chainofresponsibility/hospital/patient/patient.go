package patient

import "fmt"

type Patient struct {
	Name              string
	RegistrationDone  bool
	DoctorCheckupDone bool
	MedicationDone    bool
	PaymentDone       bool
}

func (p Patient) String() string {
	return fmt.Sprintf("Name:%s, Registration:%v, DoctorCheckup:%v, Medication:%v, Payment:%v",
		p.Name, p.RegistrationDone, p.DoctorCheckupDone, p.MedicationDone, p.PaymentDone)
}
