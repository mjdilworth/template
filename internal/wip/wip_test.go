package wip

import (
	"fmt"
	"strings"
	"testing"
)

func TestBreaker(t *testing.T) {

	app := New("foo")

	//retFunction := app.Breaker()

	fmt.Println(*app)

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

	app := New("foo")
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

	//sub test
	t.Run("simple single test give 2 and want 4", func(t *testing.T) {
		got := app.Two(2)
		want := 4

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

	})

	t.Run("simple table tests give 2 and want ...", func(t *testing.T) {
		var tests = []struct {
			name  string
			input int
			want  int
		}{
			// the table itself
			{"give 9 and want 18 ", 9, 18},
			{"give 3 and want 6", 3, 6},
			{"give 1 and want 2", 1, 2},
			{"give 0 and want 0", 0, 0},
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
	})
}
