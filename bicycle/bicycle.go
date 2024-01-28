package bicycle

import (
	"fmt"
)

const (
	DefaultBrand  = "DefaultBicycleBrand"
	DefaultColor  = "DefaultBicycleColor"
	DefaultWeight = 24
)

type Bicycle struct {
	Brand  string
	Color  string
	Weight int
}

func NewBicycle(properties ...func(*Bicycle)) *Bicycle {
	bike := &Bicycle{
		Brand:  DefaultBrand,
		Color:  DefaultColor,
		Weight: DefaultWeight,
	}

	// Apply specified properties
	for _, property := range properties {
		property(bike)
	}

	return bike
}

func WithBrand(brand string) func(*Bicycle) {
	return func(b *Bicycle) {
		b.Brand = brand
	}
}

func WithColor(color string) func(*Bicycle) {
	return func(b *Bicycle) {
		b.Color = color
	}
}

func WithWeight(weight int) func(*Bicycle) {
	return func(b *Bicycle) {
		b.Weight = weight
	}
}

func (b *Bicycle) Pedal() {
	fmt.Println("Pedaling the bicycle.")
}

func (b *Bicycle) Brake() {
	fmt.Println("Applying brakes to the bicycle.")
}
