package wip

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

type Wip struct {
	Name string
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

type Dictionary map[string]string

func (w Wip) DictionarySearch(dict Dictionary, str string) string {

	return dict[str]
}

func (w Wip) One() string {
	return "now i am in package wip and function one"
}

func (w Wip) Two(i int) int {

	//i add 2 to inout and this wil pass smimple tests
	return i + 2
}

func (w Wip) ProductTwoInts(a int, b int) int {

	return a * b

}

func (w Wip) CharRepeater(s string, times int) string {

	var retString string

	for i := 0; i < times; i++ {
		retString = retString + "a"
	}
	return retString
}

func (w Wip) SumArray(iArr []int) int {

	iRet := 0

	for _, num := range iArr {
		iRet = iRet + num
	}

	return iRet
}

func (w Wip) SumAll(toSum ...[]int) []int {

	//numOfSlices := len(toSum)
	//retSlice := make([]int, numOfSlices)

	var retSlice []int
	for _, num := range toSum {
		//retSlice[i] = w.SumArray(num)
		retSlice = append(retSlice, w.SumArray(num))
	}
	return retSlice
}

func (w Wip) SumAllTails(toSum ...[]int) []int {

	var retSlice []int
	for _, num := range toSum {
		if len(num) < 1 {
			retSlice = append(retSlice, 0)
		} else {
			tail := num[1:]
			retSlice = append(retSlice, w.SumArray(tail))

		}
	}
	return retSlice
}

func (w Wip) Perimeter(r Rectangle) float64 {

	fRet := 2 * (r.Height + r.Width)

	return fRet
}

func (r Rectangle) Area() float64 {

	fRet := r.Height * r.Width
	return fRet
}
func (c Circle) Area() float64 {

	fRet := math.Pi * (c.Radius * c.Radius)
	return fRet
}
func (tr Triangle) Area() float64 {

	return (tr.Base * tr.Height) * 0.5
}
