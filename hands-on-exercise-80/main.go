package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	user1 := User{
		Name: "RK",
		Age:  25,
	}
	user2 := User{
		Name: "Jimmy",
		Age:  30,
	}
	users := []User{user1, user2}

	fmt.Println(users)
	err := json.NewEncoder(os.Stdout).Encode(&users)
	if err != nil {
		fmt.Println(err)
	}
}
