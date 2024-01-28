package golf

import (
	"testing"
)

func TestGolfClubWithSingleProperty(t *testing.T) {
	expected := GolfClub{
		Brand:  "Titleist",
		Type:   "DefaultType",
		Length: 999,
	}

	club := NewGolfClubBuilder().
		Brand("Titleist").
		Build()

	assertGolfClubProperties(t, expected, *club, "Golf Club with single property")
}

func TestGolfClubWithNoProperties(t *testing.T) {
	expected := GolfClub{
		Brand:  "DefaultBrand",
		Type:   "DefaultType",
		Length: 999,
	}

	club := NewGolfClubBuilder().
		Build()

	assertGolfClubProperties(t, expected, *club, "Golf Club with no properties")
}

func TestGolfClubWithAllProperties(t *testing.T) {
	expected := GolfClub{
		Brand:  "Titleist",
		Type:   "Putter",
		Length: 35,
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
