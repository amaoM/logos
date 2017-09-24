package parallel

import "testing"

func BenchmarkLogos1(b *testing.B) {
	logos1()
}

func BenchmarkLogos2(b *testing.B) {
	logos2()
}

func BenchmarkLogos3(b *testing.B) {
	logos3()
}

func BenchmarkLogos4(b *testing.B) {
	logos4()
}
