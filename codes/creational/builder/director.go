package builder

type director struct {
	builder iBuilder
}

func newDirector(builder iBuilder) *director {
	return &director{
		builder: builder,
	}
}

func (d *director) setBuilder(b iBuilder) {
	d.builder = b
}

func (d *director) buildHouse() house {
	d.builder.SetDoorType()
	d.builder.SetWindowType()
	d.builder.SetNumFloor()
	return d.builder.GetHouse()
}
