package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	cetaknama()
}

func cetaknama() {
	var friend = []*Person{
		{name: "Frisky", age: 20},
		{name: "Yoseph Bram", age: 30},
		{name: "Rizky", age: 17},
		{name: "Tony", age: 40},
		{name: "Dina", age: 27},
		{name: "Aimand", age: 31},
		{name: "Devi", age: 33},
		{name: "Ellen", age: 21},
		{name: "Yusuf", age: 19},
		{name: "Meuthia", age: 18},
	}

	var nama = func(name []*Person) {
		for _, x := range name {
			fmt.Println(x.name, x.age)

		}
	}
	nama(friend)
}
