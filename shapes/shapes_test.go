package shapes

import (
	"fmt"
	"testing"
)

func TestPerimeter(t *testing.T) {
	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}

	t.Run("perimeter", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := rectangle.Perimeter()
		want := 40.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})

	t.Run("area", func(t *testing.T) {
		areaTests := []struct {
			shape Shape
			want  float64
		}{
			{Rectangle{12, 6}, 72.0},
			{Circle{10}, 314.1592653589793},
			{Triangle{12, 6}, 36.0},
		}

		for _, tt := range areaTests {

			checkArea(t, tt.shape, tt.want)

		}
	})
}

func ExampleRectangle_Area() {
	rectangle := Rectangle{12, 6}
	area := rectangle.Area()
	fmt.Println(area)
	// Output: 72
}
