package command

import (
	"context"
	"fmt"
)

type command struct{}

func New() command {
	return command{}
}

// run the command line and do what was asked
func (c command) Run(ctx context.Context) {

	//parse commnad line options do required work and exit

	//for mow i will just print stuff

	fmt.Printf("what is in our context : %v\n", ctx)
	fmt.Println("Executing run of the command line option - this will do somehting and then exit")
	//fmt.Println(ctx)

}
