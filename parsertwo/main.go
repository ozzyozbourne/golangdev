package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type permissions map[string]bool

type user struct {
	Name        string      `json:"userName"`
	Password    string      `json:"-"`
	Permissions permissions `json:"perm,omitempty"`
}

func main() {
	users := []user{{"osaid", "auto", nil},
		{"god", "42", permissions{"admin": true}},
		{"devil", "666", permissions{"admin": false, "user_access": true}}}

	//	fmt.Printf("%+v\n", users)

	out, err := json.MarshalIndent(users, "", "\t")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%s\n", out)

}
