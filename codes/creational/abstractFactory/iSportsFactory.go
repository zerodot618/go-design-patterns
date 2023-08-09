package abstractfactory

import "fmt"

type iSportsFactory interface {
	makeShoe() iShoe
	makeShort() iShort
}

func getSportsFactory(brand string) (iSportsFactory, error) {
	if brand == "adidas" {
		return &adidasFactory{}, nil
	}
	if brand == "nike" {
		return &nikeFactory{}, nil
	}
	return nil, fmt.Errorf("Wrong brand type passed")
}
