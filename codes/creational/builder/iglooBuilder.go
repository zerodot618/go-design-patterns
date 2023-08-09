package builder

type iglooBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newIglooBuilder() *iglooBuilder {
	return &iglooBuilder{}
}

func (b *iglooBuilder) SetWindowType() {
	b.windowType = "Snow Window"
}

func (b *iglooBuilder) SetDoorType() {
	b.doorType = "Snow Door"
}

func (b *iglooBuilder) SetNumFloor() {
	b.floor = 1
}

func (b *iglooBuilder) GetHouse() house {
	return house{
		doorType:   b.doorType,
		windowType: b.windowType,
		floor:      b.floor,
	}
}
