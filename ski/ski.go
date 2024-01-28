package ski

import (
	"fmt"
)

type Ski struct {
	Brand  string
	Color  string
	Length int
}

func NewSkiPrototype() *Ski {
	return &Ski{
		Brand:  "DefaultBrand",
		Color:  "DefaultColor",
		Length: 170,
	}
}

func (s *Ski) Clone() *Ski {
	return &Ski{
		Brand:  s.Brand,
		Color:  s.Color,
		Length: s.Length,
	}
}

func (s *Ski) SkiDownhill() {
	fmt.Println("Skiing downhill.")
}

func (s *Ski) Stop() {
	fmt.Println("Stopping on skis.")
}
