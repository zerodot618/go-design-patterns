package decorator

type tomatoTopping struct {
	pizza pizza
}

func (t *tomatoTopping) getPrice() int {
	return t.pizza.getPrice() + 10
}
