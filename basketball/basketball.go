package basketball

import "fmt"

// Ball represents a basketball.
type Ball struct {
	Size   int
	Color  string
	Weight int
}

// BallFactory is the interface for creating basketballs.
type BallFactory interface {
	CreateBall() *Ball
}

// NewBasketballFactory creates a factory for creating basketballs.
func NewBasketballFactory() BallFactory {
	return &BasketballFactory{}
}

// BasketballFactory is the concrete factory for creating basketballs.
type BasketballFactory struct{}

// CreateBall creates a basketball with default properties.
func (f *BasketballFactory) CreateBall() *Ball {
	return &Ball{
		Size:   7, // Standard basketball size
		Color:  "Orange",
		Weight: 22, // Standard basketball weight in ounces
	}
}

// WithSize is a factory method option to set the size of the basketball.
func WithSize(size int) func(*Ball) {
	return func(b *Ball) {
		b.Size = size
	}
}

// WithColor is a factory method option to set the color of the basketball.
func WithColor(color string) func(*Ball) {
	return func(b *Ball) {
		b.Color = color
	}
}

// WithWeight is a factory method option to set the weight of the basketball.
func WithWeight(weight int) func(*Ball) {
	return func(b *Ball) {
		b.Weight = weight
	}
}

