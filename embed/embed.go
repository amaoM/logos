package embed

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
func (human *Human) Talk() string {
	return "I am a " + human.Type
}

func logos1() string {
	man := Man{&Human{Type: "man"}}
	return man.Talk()
}
