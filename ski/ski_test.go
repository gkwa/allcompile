package ski

import (
	"reflect"
	"testing"
)

func equalSki(a, b *Ski) bool {
	return reflect.DeepEqual(a, b)
}

func TestSkiPrototype(t *testing.T) {
	testCases := []struct {
		desc     string
		modifyFn func(*Ski)
		expected *Ski
	}{
		{
			desc: "Ski 1: Customized properties",
			modifyFn: func(s *Ski) {
				s.Brand = "Atomic"
				s.Color = "Red"
				s.Length = 180
			},
			expected: &Ski{
				Brand:  "Atomic",
				Color:  "Red",
				Length: 180,
			},
		},
		{
			desc: "Ski 2: Different Color",
			modifyFn: func(s *Ski) {
				s.Color = "Blue"
			},
			expected: &Ski{
				Brand:  "DefaultBrand",
				Color:  "Blue",
				Length: 170,
			},
		},
		{
			desc: "Ski 3: Default Prototype",
			modifyFn: func(s *Ski) {
				// No modifications, using default prototype
			},
			expected: &Ski{
				Brand:  "DefaultBrand",
				Color:  "DefaultColor",
				Length: 170,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			prototype := NewSkiPrototype()
			tc.modifyFn(prototype)
			if !equalSki(prototype, tc.expected) {
				t.Errorf("Expected %+v, Got %+v", tc.expected, prototype)
			}
		})
	}
}
