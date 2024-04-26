package jsonpractice

import (
	"encoding/json"
	"fmt"
)

type Ninja struct {
	Name   string `json:"first_name"`
	Weapon string `json:"weapon_name"`
	Level  int    `json:"player_level"`
}

func JsonMain() {
	rawJson := `{
		"first_name":"Sefat Anam",
		"weapon_name":"Sword",
		"player_level":1
	}`

	var wallace Ninja

	err := json.Unmarshal([]byte(rawJson), &wallace)

	if err != nil {
		fmt.Println("Can't able to unmarshal json.")
	}

	fmt.Printf("The converted valuen struct is : %+v\n", wallace)

	cTo, err := json.Marshal(wallace)
	if err != nil {
		fmt.Println("Can't able to marshal json.")
	}
	fmt.Printf("The converted json value is : %+s\n", cTo)

}
