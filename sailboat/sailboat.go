package sailboat

type Sailboat struct {
	Brand  string
	Color  string
	Length int
}

// NewSailboat creates a new Sailboat instance with default values and allows specifying properties.
func NewSailboat(properties ...func(*Sailboat)) *Sailboat {
	boat := &Sailboat{
		Brand:  "Unknown",
		Color:  "Unspecified",
		Length: 999,
	}

	// Apply specified properties
	for _, property := range properties {
		property(boat)
	}

	return boat
}

// WithBrand is a function to set the brand property of a Sailboat.
func WithBrand(brand string) func(*Sailboat) {
	return func(b *Sailboat) {
		b.Brand = brand
	}
}

// WithColor is a function to set the color property of a Sailboat.
func WithColor(color string) func(*Sailboat) {
	return func(b *Sailboat) {
		b.Color = color
	}
}

// WithLength is a function to set the length property of a Sailboat.
func WithLength(length int) func(*Sailboat) {
	return func(b *Sailboat) {
		b.Length = length
	}
}
