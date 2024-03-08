package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	str := `[{"name":"RK","age":25},{"name":"Jimmy","age":30}]`

	fmt.Println(str)
	var users []User
	err := json.Unmarshal([]byte(str), &users)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(users)

}
