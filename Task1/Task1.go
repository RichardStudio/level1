package main

import "fmt"

type Human struct {
	Age  int
	Name string
}

func (h *Human) Say() {
	fmt.Printf("Human %s says i'm %d years old\n", h.Name, h.Age)
}

type Action struct {
	Human
}

func (a *Action) Walk() {
	fmt.Printf("%s is Walking\n", a.Name)
}

func main() {
	human := Action{
		Human: Human{
			Age:  15,
			Name: "James",
		},
	}

	human.Say()
	human.Walk()
}
