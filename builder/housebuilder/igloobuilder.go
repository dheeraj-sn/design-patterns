package housebuilder

type IglooBuilder struct {
	windowType string
	doorType   string
	numFloor   int
}

func NewIglooBuilder() *IglooBuilder {
	return &IglooBuilder{}
}

func (b *IglooBuilder) SetWindowType() {
	b.windowType = "Snow Window"
}
func (b *IglooBuilder) SetDoorType() {
	b.doorType = "Snow Door"
}
func (b *IglooBuilder) SetNumFloor() {
	b.numFloor = 1
}

func (b *IglooBuilder) GetHouse() House {
	return House{
		WindowType: b.windowType,
		DoorType:   b.doorType,
		Floor:      b.numFloor,
		Type:       "Igloo",
	}
}
