package wip

import (
	"context"
	"log/slog"
)

type Wip struct {
	Name string
}

func (w Wip) LogMe(s string, k string, v int) {
	slog.Info(s, k, v)
	// log.Fatalf("Fatal string %s, key %s and value %d", s, k, v)
}

func New(name string) *Wip {
	//if name is empty string use default value "foo"
	if len(name) < 1 {
		name = "foo"
	}
	return &Wip{
		Name: name,
	}
}

func (w Wip) One() string {
	return "now i am in package wip and function one"
}

func (w Wip) Two(i int) int {

	//i multipl 2 to inout and this wil pass smimple tests
	return (i * 2)
}

func (w Wip) Breaker(circuit func(context.Context) (string, error), failureThreshold uint) func(context.Context) (string, error) {

	//
	return nil
}

func (w Wip) myfunc(c context.Context) (string, error) {

	return "", nil
}
