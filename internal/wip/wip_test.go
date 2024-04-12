package wip

import "testing"

func TestTwo(t *testing.T) {

	app := &Wip{
		Name: "foo",
}

result := app.Two(1)

if result != 5 {

t.Errorf("Result was incorrect, got: %d, want: %d.", result, 5)
}

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
 _, tt := range tests {
t.Run(tt.name, func(t *testing.T) {
 ans:= app.Two(tt.input)
 ans != tt.want {
t.Errorf("got %d, want %d", ans, tt.want)
}
})
}
}