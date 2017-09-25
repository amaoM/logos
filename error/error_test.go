package error

import "testing"

func TestLogos1(t *testing.T) {
	err := logos1()
	if err == nil {
		t.Error("Not error")
	} else {
		t.Log(err)
	}
}

func TestLogos2(t *testing.T) {
	err := logos2()
	if err == nil {
		t.Error("Not error")
	} else {
		t.Log(err)
	}
}

func TestLogos3(t *testing.T) {
	err := logos3()
	if err == nil {
		t.Error("Not error")
	} else {
		t.Log(err)
	}
}
