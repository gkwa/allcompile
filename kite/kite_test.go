package kite

import (
	"testing"
)

func TestKiteProperties(t *testing.T) {
	testCases := []struct {
		desc        string
		constructor func() *Kite
		expected    Kite
	}{
		{
			desc: "Kite 1: all properties specified using functional options",
			constructor: func() *Kite {
				return NewKite(
					WithBrand("HighFly"),
					WithColor("Yellow"),
					WithLength(50),
				)
			},
			expected: Kite{
				Brand:  "HighFly",
				Color:  "Yellow",
				Length: 50,
			},
		},
		{
			desc: "Kite 2: single property specified using functional options",
			constructor: func() *Kite {
				return NewKite(
					WithColor("Green"),
				)
			},
			expected: Kite{
				Brand:  "DefaultBrand",
				Color:  "Green",
				Length: 999,
			},
		},
		{
			desc: "Kite 3: no properties specified, using defaults for all",
			constructor: func() *Kite {
				return NewKite()
			},
			expected: Kite{
				Brand:  "DefaultBrand",
				Color:  "DefaultColor",
				Length: 999,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			kite := tc.constructor()
			if *kite != tc.expected {
				t.Errorf("Expected %v, Got %v", tc.expected, *kite)
			}
		})
	}
}
