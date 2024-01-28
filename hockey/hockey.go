package hockey

import "fmt"

type HockeyStick struct {
	Brand  string
	Length int
	Color  string
}

func NewHockeyStick(properties ...HockeyStick) *HockeyStick {
	stick := &HockeyStick{
		Brand:  "DefaultBrand",
		Length: 60,
		Color:  "DefaultColor",
	}

	// Apply specified properties
	for _, property := range properties {
		if property.Brand != "" {
			stick.Brand = property.Brand
		}
		if property.Length != 0 {
			stick.Length = property.Length
		}
		if property.Color != "" {
			stick.Color = property.Color
		}
	}

	return stick
}

func (h *HockeyStick) Shoot() {
	fmt.Println("Shooting with the hockey stick.")
}

func (h *HockeyStick) Pass() {
	fmt.Println("Passing with the hockey stick.")
}
