package basketball

import (
	"testing"
)

func TestBasketballWithDefaultProperties(t *testing.T) {
	expected := BasketBall{
		Size:   7,
		Color:  "Orange",
		Weight: 22,
	}

	factory := NewBasketballFactory()
	basketball := factory.CreateBasketBall()

	assertBasketballProperties(t, expected, *basketball, "Basketball with no custom properties")
}

func TestBasketballWithCustomProperties(t *testing.T) {
	expected := BasketBall{
		Size:   6,
		Color:  "Red",
		Weight: 21,
	}

	factory := NewBasketballFactory(
		WithSize(6),
		WithColor("Red"),
		WithWeight(21),
	)

	basketball := factory.CreateBasketBall()
	assertBasketballProperties(t, expected, *basketball, "Basketball all properties explicitly set using With* functions")
}

func TestBasketballWithSingleCustomProperty(t *testing.T) {
	expected := BasketBall{
		Size:   11,
		Color:  "Orange",
		Weight: 22,
	}

	factory := NewBasketballFactory(
		WithSize(11),
	)

	basketball := factory.CreateBasketBall()
	assertBasketballProperties(t, expected, *basketball, "Basketball with single custom property explicitly set using With* functions")
}

func assertBasketballProperties(t *testing.T, expected, actual BasketBall, desc string) {
	t.Run(desc, func(t *testing.T) {
		if expected != actual {
			t.Errorf("Expected %v, Got %v", expected, actual)
		}
	})
}
