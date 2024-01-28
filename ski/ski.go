package ski

import (
	"fmt"
)

type Ski struct {
	Brand  string
	Color  string
	Weight int
}

const (
	DefaultBrand  = "DefaultSkiBrand"
	DefaultColor  = "DefaultSkiColor"
	DefaultWeight = 170
)

func NewSkiPrototype() *Ski {
	return &Ski{
		Brand:  DefaultBrand,
		Color:  DefaultColor,
		Weight: DefaultWeight,
	}
}

func (s *Ski) Clone() *Ski {
	return &Ski{
		Brand:  s.Brand,
		Color:  s.Color,
		Weight: s.Weight,
	}
}

func (s *Ski) SkiDownhill() {
	fmt.Println("Skiing downhill.")
}

func (s *Ski) Stop() {
	fmt.Println("Stopping on skis.")
}
