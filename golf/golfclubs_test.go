package golf

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

func TestUnmarshalGolfClubsFromJSON(t *testing.T) {
	// JSON representation of a slice of golf clubs
	jsonData := `[
		{
			"Brand": "Titleist",
			"Type": "Driver",
			"Length": 45
		},
		{
			"Brand": "Callaway",
			"Type": "Iron",
			"Length": 36
		}
	]`

	var clubs GolfClubs
	err := json.Unmarshal([]byte(jsonData), &clubs)

	expected := GolfClubs{
		&GolfClub{Brand: "Titleist", Type: "Driver", Length: 45},
		&GolfClub{Brand: "Callaway", Type: "Iron", Length: 36},
	}

	if err != nil {
		t.Errorf("Error unmarshaling: %v", err)
	} else if !reflect.DeepEqual(clubs, expected) {
		t.Errorf("Expected %v, Got %v", expected, clubs)
	} else {
		prettyJSON, _ := json.MarshalIndent(clubs, "", "  ")
		fmt.Println("Pretty Printed JSON:")
		fmt.Println(string(prettyJSON))
	}
}

func TestNewGolfClubsAndAddClub(t *testing.T) {
	clubs := NewGolfClubs()

	club1 := &GolfClub{Brand: "Titleist", Type: "Driver", Length: 45}
	clubs = clubs.AddClub(club1)

	club2 := &GolfClub{Brand: "Callaway", Type: "Iron", Length: 36}
	clubs = clubs.AddClub(club2)

	expected := GolfClubs{club1, club2}

	if !reflect.DeepEqual(clubs, expected) {
		t.Errorf("Expected %v, Got %v", expected, clubs)
	}
}
