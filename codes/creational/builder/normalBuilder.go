package builder

type normalBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newNormalBuilder() *normalBuilder {
	return &normalBuilder{}
}

func (n *normalBuilder) SetWindowType() {
	n.windowType = "Wooden Window"
}

func (n *normalBuilder) SetDoorType() {
	n.doorType = "Wooden Door"
}

func (n *normalBuilder) SetNumFloor() {
	n.floor = 2
}

func (b *normalBuilder) GetHouse() house {
	return house{
		doorType:   b.doorType,
		windowType: b.windowType,
		floor:      b.floor,
	}
}
