package golf

import (
	"testing"
)

func TestGolfClubWithSingleCustomizedProperty(t *testing.T) {
	expected := GolfClub{
		Brand:  "Titleist",
		Color:  DefaultColor,
		Weight: DefaultWeight,
	}

	club := NewGolfClubBuilder().
		Brand("Titleist").
		Build()

	assertGolfClubProperties(t, expected, *club, "Golf Club with single property")
}

func TestGolfClubWithNoCustomizedProperties(t *testing.T) {
	expected := GolfClub{
		Brand:  DefaultBrand,
		Color:  DefaultColor,
		Weight: DefaultWeight,
	}

	club := NewGolfClubBuilder().
		Build()

	assertGolfClubProperties(t, expected, *club, "Golf Club with no properties")
}

func TestGolfClubWithAllCustomProperties(t *testing.T) {
	expected := GolfClub{
		Brand:  "Titleist",
		Color:  "Putter",
		Weight: 35,
	}

	club := NewGolfClubBuilder().
		Type("Putter").
		Brand("Titleist").
		Length(35).
		Build()

	assertGolfClubProperties(t, expected, *club, "Golf Club with all properties")
}

func assertGolfClubProperties(t *testing.T, expected, actual GolfClub, desc string) {
	t.Run(desc, func(t *testing.T) {
		if expected != actual {
			t.Errorf("Expected %v, Got %v", expected, actual)
		}
	})
}
