package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string `json:"name"`
	Last string `json:"last"`
	Age  int    `json:"age"`
}

type ByAge []User

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type ByLast []User

func (l ByLast) Len() int           { return len(l) }
func (l ByLast) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l ByLast) Less(i, j int) bool { return l[i].Last < l[j].Last }

func main() {
	user1 := User{
		Name: "RK",
		Last: "Kumar",
		Age:  25,
	}
	user2 := User{
		Name: "Sam",
		Last: "Joe",
		Age:  35,
	}
	user3 := User{
		Name: "Jimmy",
		Last: "Doe",
		Age:  30,
	}
	users := []User{user1, user2, user3}

	fmt.Println("Before sorting", users)

	sort.Sort(ByAge(users))
	fmt.Println("Sorting using age :", users)

	sort.Sort(ByLast(users))
	fmt.Println("Sorting using last :", users)

}
