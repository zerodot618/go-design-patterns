package decorator

type veggeMania struct {
}

func (m *veggeMania) getPrice() int {
	return 15
}
