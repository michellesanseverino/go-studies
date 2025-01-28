package main

type Dog struct {
	Race  string
	Color string
	Paws  int
}

type Cat struct {
	Color string
	Paws  int
}

func (d Dog) Noise() string {
	return "Woof!"
}

func (c Cat) Noise() string {
	return "Meow!"
}

func (d Dog) NumberOfPaws() int {
	return d.Paws
}

func (c Cat) NumberOfPaws() int {
	return c.Paws
}

// Animal interface
type Animal interface {
	Noise() string
	NumberOfPaws() int
}

func MakeNoise(a Animal) string {
	return a.Noise()
}

func NumberOfPaws(a Animal) int {
	return a.NumberOfPaws()
}

func main() {
	dog1 := Dog{"Husky", "White", 4}
	cat1 := Cat{"Black", 4}

	MakeNoise(dog1)
	MakeNoise(cat1)

}
