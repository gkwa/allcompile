package bicycle

import (
	"fmt"
	"strings"
)

type Bicycle struct {
	Brand  string
	Color  string
	Weight int
}

func (b *Bicycle) String() string {
	var builder strings.Builder

	builder.WriteString("Brand: ")
	builder.WriteString(b.Brand)
	builder.WriteString("\nColor: ")
	builder.WriteString(b.Color)
	builder.WriteString("\nWeight: ")
	builder.WriteString(fmt.Sprintf("%d", b.Weight))

	return builder.String()
}

func (b *Bicycle) Pedal() {
	fmt.Println("Pedaling the bicycle.")
}

func (b *Bicycle) Brake() {
	fmt.Println("Applying brakes to the bicycle.")
}

// NewBicycle creates a new Bicycle instance with default values.
func NewBicycle() *Bicycle {
	return &Bicycle{
		Brand:  "Unknown",
		Color:  "Unspecified",
		Weight: 0,
	}
}
