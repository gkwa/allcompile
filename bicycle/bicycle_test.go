package bicycle

import (
	"testing"
)

func TestBicycleProperties(t *testing.T) {
	testCases := []struct {
		desc        string
		constructor func() *Bicycle
		expected    Bicycle
	}{
		{
			desc: "Bicycle 1: all properties specified using functional options",
			constructor: func() *Bicycle {
				return NewBicycle(
					WithBrand("Giant"),
					WithColor("Blue"),
					WithWeight(10),
				)
			},
			expected: Bicycle{
				Brand:  "Giant",
				Color:  "Blue",
				Weight: 10,
			},
		},
		{
			desc: "Bicycle 2: single property specified using functional options",
			constructor: func() *Bicycle {
				return NewBicycle(
					WithColor("Green"),
				)
			},
			expected: Bicycle{
				Brand:  "DefaultBrand",
				Color:  "Green",
				Weight: 999,
			},
		},
		{
			desc: "Bicycle 3: composite literal with all properties specified",
			constructor: func() *Bicycle {
				bike := &Bicycle{
					Brand:  "Specialized",
					Color:  "Red",
					Weight: 12,
				}
				bike.Brake()
				return bike
			},
			expected: Bicycle{
				Brand:  "Specialized",
				Color:  "Red",
				Weight: 12,
			},
		},
		{
			desc: "Bicycle 4: composite literal single property specified",
			constructor: func() *Bicycle {
				bike := &Bicycle{
					Brand: "SpeedX",
				}
				bike.Pedal()
				return bike
			},
			expected: Bicycle{
				Brand:  "SpeedX",
				Color:  "", // Expecting an empty string as the zero value
				Weight: 0,
			},
		},
		{
			desc: "Bicycle 5: two properties specified using variadic constuctor",
			constructor: func() *Bicycle {
				bike := NewBicycle(
					WithBrand("Giant"),
					WithWeight(10),
				)
				bike.Pedal()
				return bike
			},
			expected: Bicycle{
				Brand:  "Giant",
				Color:  "DefaultColor",
				Weight: 10,
			},
		},
		{
			desc: "Bicycle 6: no properties specified, using defaults for all",
			constructor: func() *Bicycle {
				bike := NewBicycle()
				bike.Pedal()
				return bike
			},
			expected: Bicycle{
				Brand:  "DefaultBrand",
				Color:  "DefaultColor",
				Weight: 999,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			bike := tc.constructor()
			if *bike != tc.expected {
				t.Errorf("Expected %v, Got %v", tc.expected, *bike)
			}
		})
	}
}
