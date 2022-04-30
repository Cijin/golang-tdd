package main

import "testing"

func TestPerimeter(t *testing.T) {
	r := Rectange{10.0, 10.0}
	got := perimeter(r)
	want := 40.0

	if got != want {
		t.Errorf("Got %g want %g", got, want)
	}
}

func TestArea(t *testing.T) {

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()

		if got != want {
			t.Errorf("%T: Got %g want %g", shape, got, want)
		}
	}

	areaTests := []struct {
		s    Shape
		want float64
	}{
		{s: Rectange{W: 10.0, H: 10.0}, want: 100.00},
		{s: Circle{R: 10.0}, want: 314.1592653589793},
		{s: Triangle{H: 12.0, B: 6.0}, want: 36.0},
	}

	for _, tt := range areaTests {
		checkArea(t, tt.s, tt.want)
	}
}
