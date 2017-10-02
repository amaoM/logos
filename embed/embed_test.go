package embed

import "testing"

func TestLogos1(t *testing.T) {
	expected := "logos1Main"
	actual := logos1()
	if expected != actual {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestLogos2(t *testing.T) {
	expected := "logos2InterfaceMain"
	actual := logos2()
	if expected != actual {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestLogos3(t *testing.T) {
	expected := "logos2InterfaceMain"
	actual := logos3()
	if expected != actual {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}
