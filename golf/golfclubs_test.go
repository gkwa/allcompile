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
			"Color": "Purple",
			"Weight": 45
		},
		{
			"Brand": "Callaway",
			"Color": "Brown",
			"Weight": 36
		}
	]`

	var clubs GolfClubs
	err := json.Unmarshal([]byte(jsonData), &clubs)

	expected := GolfClubs{
		&GolfClub{Brand: "Titleist", Color: "Purple", Weight: 45},
		&GolfClub{Brand: "Callaway", Color: "Brown", Weight: 36},
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
	var clubs GolfClubs

	club1 := &GolfClub{Brand: "Titleist", Color: "Blue", Weight: 45}
	clubs = clubs.AddClub(club1)

	club2 := &GolfClub{Brand: "Callaway", Color: "Maroon", Weight: 36}
	clubs = clubs.AddClub(club2)

	expected := GolfClubs{club1, club2}

	if !reflect.DeepEqual(clubs, expected) {
		t.Errorf("Expected %v, Got %v", expected, clubs)
	}
}
