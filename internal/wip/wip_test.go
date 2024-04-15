package wip

import (
	"strings"
	"testing"
)



func TestNewWip(t *testing.T) {
	//not much of a test really
	instance := NewWip("")
	//check that
	if len(instance.Name) < 1 {
		t.Errorf("Missing value for field: " + instance.Name)
	}

}
func TestOne(t *testing.T) {
	app := &Wip{
		Name: "foo",
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
