package entity

type Animal struct {
	ID    int
	Name  string
	Class string
	Legs  int
}

type CreateAnimal struct {
	Name  string
	Class string
	Legs  int
}
