package main

import (
	"encoding/json"
	"fmt"
	"log"
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
	bs, err := convertIntoArrByte(users)
	if err != nil {
		log.Fatalln(err)
		return
	}

	fmt.Println(string(bs))

}

func convertIntoArrByte(users []User) ([]byte, error) {
	bs, err := json.Marshal(users)
	if err != nil {
		return []byte{}, fmt.Errorf("Error while marshalling : %v", err)
	}
	return bs, nil
}
