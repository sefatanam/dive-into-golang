package builder

import (
	"errors"
	"fmt"
)

type Burger struct {
	size      int
	cheese    string
	pepperoni bool
	tomato    bool
}

type BurgerBuilder interface {
	AddSize(n int) *Burger
	AddPepperoni() *Burger
	AddCheese(n string) *Burger
	AddTomato() *Burger
}

func NewBurger() BurgerBuilder {
	return &Burger{}
}
func (b *Burger) AddSize(size int) *Burger {
	b.size = size
	return b
}

func (b *Burger) AddPepperoni() *Burger {
	b.pepperoni = true
	return b
}

func (b *Burger) AddCheese(name string) *Burger {
	b.cheese = name
	return b
}

func (b *Burger) AddTomato() *Burger {
	b.tomato = true
	return b
}

func (b *Burger) Build() (*Burger, error) {
	if b.size <= 0 {
		return nil, errors.New("burger size is not defined")
	}
	return b, nil
}

func RunBuilder() {
	burger, err := NewBurger().
		AddCheese("White").
		AddSize(12).
		AddPepperoni().
		Build()

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(burger)
}
