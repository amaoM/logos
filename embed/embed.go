package embed

import "fmt"

// Human ...
type Human struct {
	Type string
}

// Man ...
type Man struct {
	*Human
}

// Woman ...
type Woman struct {
	*Human
}

// Talk ...
func (human *Human) Talk() bool {
	fmt.Println("I am a " + human.Type)
	return true
}

func logos1() bool {
	man := Man{&Human{Type: "man"}}
	return man.Talk()
}

type Animal interface {
	Eat()
}

type Dog struct {
}

type Cat struct {
}

func (d *Dog) Eat() bool {
	fmt.Println("I eating dog foods.")
	return true
}

func (c *Cat) Eat() bool {
	fmt.Println("I eating cat foods.")
	return true
}

func (h *Human) Eat() bool {
	fmt.Println("I eating tenpura")
	return true
}

func (m *Man) Ean() bool {
	fmt.Println("I eating steaks")
	return true
}

func (wm *Woman) Eat() bool {
	fmt.Println("I eating cakes")
	return true
}

func logos2() bool {
	dog := new(Dog)
	cat := new(Cat)
	human := new(Human)
	man := new(Man)
	woman := new(Woman)
	dog.Eat()
	cat.Eat()
	human.Eat()
	man.Eat()
	woman.Eat()

	return true
}
