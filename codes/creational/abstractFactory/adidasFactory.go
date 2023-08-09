package abstractfactory

type adidasFactory struct{}

func (a *adidasFactory) makeShoe() iShoe {
	return &adidasShoe{
		shoe: shoe{
			logo: "adidas",
			size: 14,
		},
	}
}

func (a *adidasFactory) makeShort() iShort {
	return &adidasShort{
		short: short{
			logo: "adidas",
			size: 14,
		},
	}
}
