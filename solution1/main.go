package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

func (h *Human) describe() {
	fmt.Printf("Hi, my name is %s\n", h.Name)
}

type Action struct {
	Act   string
	Human // embedded struct
}

func main() {
	var action Action

	// Initialization struct fields.
	action.Act = "progrmming"
	action.Name = "Amir"
	action.Age = 19

	// Using inherited method.
	action.describe()
	fmt.Printf("%s is %d years old and he is %s\n", action.Name, action.Age, action.Act)
}
