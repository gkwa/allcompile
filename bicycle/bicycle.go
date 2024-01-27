package bicycle

import (
	"fmt"
)

type Bicycle struct {
	Brand  string
	Color  string
	Weight int
}

// NewBicycle creates a new Bicycle instance with default values and allows specifying properties.
func NewBicycle(properties ...func(*Bicycle)) *Bicycle {
	bike := &Bicycle{
		Brand:  "Unknown",
		Color:  "Unspecified",
		Weight: 0,
	}

	// Apply specified properties
	for _, property := range properties {
		property(bike)
	}

	return bike
}

// WithBrand is a function to set the brand property of a Bicycle.
func WithBrand(brand string) func(*Bicycle) {
	return func(b *Bicycle) {
		b.Brand = brand
	}
}

// WithColor is a function to set the color property of a Bicycle.
func WithColor(color string) func(*Bicycle) {
	return func(b *Bicycle) {
		b.Color = color
	}
}

// WithWeight is a function to set the weight property of a Bicycle.
func WithWeight(weight int) func(*Bicycle) {
	return func(b *Bicycle) {
		b.Weight = weight
	}
}

// Pedal peddles the bicycle.
func (b *Bicycle) Pedal() {
	fmt.Println("Pedaling the bicycle.")
}

// Brake applies the brakes to the bicycle.
func (b *Bicycle) Brake() {
	fmt.Println("Applying brakes to the bicycle.")
}
