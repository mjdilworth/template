package wip

import (
	"fmt"
)

type Wip struct {
	Name string
}

func NewWip(name string) *Wip {
	return &Wip{
		Name: name,
	}
}

func (w Wip) One() {
	fmt.Println("now i am in package wip and function one")

}

func (w Wip) Two(i int) int {

	return i + 3
}
