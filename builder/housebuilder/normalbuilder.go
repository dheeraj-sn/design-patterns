package housebuilder

type NormalBuilder struct {
	windowType string
	doorType   string
	numFloor   int
}

func NewNormalBuilder() *NormalBuilder {
	return &NormalBuilder{}
}

func (b *NormalBuilder) SetWindowType() {
	b.windowType = "Wooden Window"
}
func (b *NormalBuilder) SetDoorType() {
	b.doorType = "Wooden Door"
}
func (b *NormalBuilder) SetNumFloor() {
	b.numFloor = 2
}

func (b *NormalBuilder) GetHouse() House {
	return House{
		WindowType: b.windowType,
		DoorType:   b.doorType,
		Floor:      b.numFloor,
		Type:       "Normal House",
	}
}
