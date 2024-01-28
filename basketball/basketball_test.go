package basketball

import (
	"testing"
)

func TestBasketballProperties(t *testing.T) {
	testCases := []struct {
		desc        string
		constructor func() BallFactory
		expected    BasketBall
	}{
		{
			desc: "Basketball 1: all properties specified using functional options",
			constructor: func() BallFactory {
				return NewBasketballFactory()
			},
			expected: BasketBall{
				Size:   7,
				Color:  "Orange",
				Weight: 22,
			},
		},
		{
			desc: "Basketball 2: custom properties specified using functional options",
			constructor: func() BallFactory {
				return NewBasketballFactory(
					WithSize(6),
					WithColor("Red"),
					WithWeight(21),
				)
			},
			expected: BasketBall{
				Size:   6,
				Color:  "Red",
				Weight: 21,
			},
		},
		{
			desc: "Basketball 3: single custom property specified using functional options",
			constructor: func() BallFactory {
				return NewBasketballFactory(
					WithSize(11),
				)
			},
			expected: BasketBall{
				Size:   11,
				Color:  "Orange",
				Weight: 22,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			ballFactory := tc.constructor()
			basketball := ballFactory.CreateBasketBall()
			if *basketball != tc.expected {
				t.Errorf("Expected %v, Got %v", tc.expected, *basketball)
			}
		})
	}
}
