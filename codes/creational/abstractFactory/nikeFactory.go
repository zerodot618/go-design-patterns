package abstractfactory

type nikeFactory struct{}

func (n *nikeFactory) makeShoe() iShoe {
	return &nikeShoe{
		shoe: shoe{
			logo: "nike",
			size: 14,
		},
	}
}

func (n *nikeFactory) makeShort() iShort {
	return &nikeShort{
		short: short{
			logo: "nike",
			size: 14,
		},
	}
}
