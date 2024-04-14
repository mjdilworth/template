package wip

type Wip struct {
	Name string
}

func NewWip(name string) *Wip {
	//if name is empty string use default value "foo"
	if len(name) <1 {
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

	//i add 2 to inout and this wil pass smimple tests
	return i + 2
}
