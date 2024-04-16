package wip

import (
	"reflect"
	"strings"
	"testing"
)

func TestArea(t *testing.T) {

	areaTests := []struct {
		name  string
		shape Shape

		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{name: "circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "triangle", shape: Triangle{Height: 12, Base: 6}, hasArea: 36.00},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.hasArea {
			t.Errorf("got %g want %g", got, tt.hasArea)
		}
	}
}

func TestPerimeter(t *testing.T) {
	app := &Wip{"foo"}

	rectangle := Rectangle{10.00, 10.00}
	got := app.Perimeter(rectangle)

	want := 40.00

	if got != want {
		t.Errorf("wanted '%0.2f' but got '%0.2f'", want, got)
	}

}

func TestSumAllTails(t *testing.T) {
	app := &Wip{"foo"}

	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("wanted '%v' but got '%v", want, got)
		}
	}
	t.Run("sum tails of slices", func(t *testing.T) {
		got := app.SumAllTails([]int{1, 2, 3}, []int{0, 9})
		want := []int{5, 9}
		checkSums(t, got, want)

	})
	t.Run("sum tails of slices with empty slice", func(t *testing.T) {
		got := app.SumAllTails([]int{1, 2, 3}, []int{})
		want := []int{5, 0}
		checkSums(t, got, want)

	})

}
func TestSumAll(t *testing.T) {
	app := &Wip{"foo"}

	got := app.SumAll([]int{1, 2, 3}, []int{2, 2, 2})

	want := []int{6, 6}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("wanted '%v' but got '%v", want, got)
	}

}

// to sum all numbers in an array
func TestSumArray(t *testing.T) {

	app := &Wip{"foo"}

	t.Run("collection of unknown numbers", func(t *testing.T) {
		iArray := []int{1, 2, 3, 4}
		want := 10

		got := app.SumArray(iArray)
		if got != want {
			t.Errorf("wanted '%d' but got '%d', given '%v", want, got, iArray)
		}
	})

}

// go test -bench=. from file dir
func BenchmarkCharRepeater(b *testing.B) {
	app := &Wip{"foo"}
	times := 10
	for i := 0; i < b.N; i++ {
		app.CharRepeater("a", times)
	}
}
func TestCharRepeater(t *testing.T) {

	app := &Wip{"foo"}
	times := 10
	str := "a"
	got := app.CharRepeater(str, times)
	want := strings.Repeat(str, times)
	if got != want {
		t.Errorf("wanted '%s' but got '%s'", want, got)
	}
}
func TestProductTwoInts(t *testing.T) {

	app := &Wip{
		Name: "foo",
	}

	got := app.ProductTwoInts(5, 6)
	want := 30
	if got != want {
		t.Errorf("wantedted '%d' but got '%d'", want, got)
	}
}

func TestNewWip(t *testing.T) {
	//not much of a test really
	instance := &Wip{
		Name: "bar",
	}
	//check that
	if len(instance.Name) < 1 {
		t.Errorf("Missing value for field: " + instance.Name)
	}

}
func TestOne(t *testing.T) {
	app := &Wip{
		Name: "bar",
	}
	//i can test for the presence of a string?
	//testString := app.One()
	if !strings.Contains(app.One(), "function") {
		t.Errorf("Result was incorrect, I am looking for function and, %s didnt have it", app.One())
	}
}

func TestTwo(t *testing.T) {

	app := &Wip{
		Name: "foo",
	}

	t.Run("simple single test from 1 add 2 = 3", func(t *testing.T) {
		got := app.Two(1)
		want := 3
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	var tests = []struct {
		name  string
		input int
		want  int
	}{
		// the table itself
		{"9 should be 11", 9, 11},
		{"3 should be 5", 3, 5},
		{"1 should be 3", 1, 3},
		{"0 should be 2", 0, 2},
	}
	// The execution loop
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := app.Two(tt.input)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}
