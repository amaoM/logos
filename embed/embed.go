package embed

import "log"

type logos1Main struct {
	Name string
}

type logos1Sub struct {
	logos1Main
}

func logos1() string {
	s := new(logos1Sub)
	s.Name = "logos1Main"
	return s.Name
}

type logos2Interface interface {
	Logos2Func(name string) string
}

// Hiding the logos2Interface
type logos2interface logos2Interface

// The logos2Main struct is a logos2Interface type
// because it implementate the logos2Func
type logos2Main struct{}

func (l2m *logos2Main) Logos2Func(name string) string {
	return name
}

// Embed of logos2interface
type logos2InterfaceMain struct {
	logos2interface
}

//
func newLogos2Main() *logos2InterfaceMain {
	return &logos2InterfaceMain{&logos2Main{}}
}

func logos2() string {
	l2im := newLogos2Main()
	return l2im.Logos2Func("logos2InterfaceMain")
}

func logos3() string {
	var f interface{} = newLogos2Main()
	l2i, ok := f.(logos2interface)
	if !ok {
		log.Fatal("Error")
	}
	return l2i.Logos2Func("logos2InterfaceMain")
}
