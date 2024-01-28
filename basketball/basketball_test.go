package basketball

import (
	"testing"
)

func TestBasketballWithDefaultProperties(t *testing.T) {
	expected := BasketBall{
		Brand:  DefaultBrand,
		Color:  DefaultColor,
		Weight: DefaultWeight,
	}

	factory := NewBasketballFactory()
	basketball := factory.CreateBasketBall()

	assertBasketballProperties(t, expected, *basketball, "Basketball with no custom properties")
}

func TestBasketballWithCustomProperties(t *testing.T) {
	expected := BasketBall{
		Brand:  "Wilson",
		Color:  "Red",
		Weight: 21,
	}

	factory := NewBasketballFactory(
		WithBrand("Wilson"),
		WithColor("Red"),
		WithWeight(21),
	)

	basketball := factory.CreateBasketBall()
	assertBasketballProperties(t, expected, *basketball, "Basketball all properties explicitly set using With* functions")
}

func TestBasketballWithSingleCustomProperty(t *testing.T) {
	expected := BasketBall{
		Brand:  "Tribute",
		Color:  DefaultColor,
		Weight: DefaultWeight,
	}

	factory := NewBasketballFactory(
		WithBrand("Tribute"),
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
