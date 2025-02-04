package product

import (
	"errors"
	"math/rand"
)

type Product struct {
	title string
	id    uint32
	price float32
}

func New(title string, price float32) (*Product, error) {
	if title == "" || price == 0 {
		return nil, errors.New("title and price should be present")
	}

	return &Product{
		title: title,
		price: price,
		id:    rand.Uint32(),
	}, nil
}
