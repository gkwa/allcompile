package hockey

import (
	"testing"
)

func TestHockeyStickProperties(t *testing.T) {
	testCases := []struct {
		desc        string
		constructor func() *HockeyStick
		expected    HockeyStick
	}{
		{
			desc: "HockeyStick 1: composite literal with all properties specified",
			constructor: func() *HockeyStick {
				stick := &HockeyStick{
					Brand:  "Bauer",
					Length: 63,
					Color:  "Blue",
				}
				stick.Shoot()
				return stick
			},
			expected: HockeyStick{
				Brand:  "Bauer",
				Length: 63,
				Color:  "Blue",
			},
		},
		{
			desc: "HockeyStick 2: composite literal single property specified",
			constructor: func() *HockeyStick {
				stick := &HockeyStick{
					Brand: "CCM",
				}
				stick.Pass()
				return stick
			},
			expected: HockeyStick{
				Brand:  "CCM",
				Length: 60,             // Expecting the default value
				Color:  "DefaultColor", // Expecting the default value
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			stick := tc.constructor()
			if *stick != tc.expected {
				t.Errorf("Expected %v, Got %v", tc.expected, *stick)
			}
		})
	}
}
