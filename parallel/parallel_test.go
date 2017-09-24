package parallel

import "testing"

func TestLogos1(t *testing.T) {
	if !logos1() {
		t.Error("Got false")
	}
}

func TestLogos2(t *testing.T) {
	actual := logos2()
	for _, a := range actual {
		if !a {
			t.Error("Got false")
		}
	}
}

func TestLogos3(t *testing.T) {
	actual := logos3()
	for i, a := range actual {
		if a != i+i {
			t.Errorf("%d != %d", a, i+i)
		} else {
			t.Logf("a = %d, i+i = %d", a, i+i)
		}
	}
}

func TestLogos4(t *testing.T) {
	actual := logos4()
	for i, a := range actual {
		if a != i+i {
			t.Errorf("%d != %d", a, i+i)
		} else {
			t.Logf("a = %d, i+i = %d", a, i+i)
		}
	}
}

func TestLogos5(t *testing.T) {
	actual := logos5()
	if actual != 100000 {
		t.Log("Not thread(goroutine) safe")
	}
}

func TestLogos6(t *testing.T) {
	actual := logos6()
	if actual != 100000 {
		t.Errorf("Not thread(goroutine) safe")
	}
}

func TestLogos7(t *testing.T) {
	actual := logos7()
	if actual != 100000 {
		t.Errorf("Not thread(goroutine) safe")
	}
}
