package main

import (
	"flag"
	"fmt"

	"github.com/mjdilworth/template/internal/wip"
)

var Version = "development"
var CommitId string
var BuildTimestamp string

func main() {

	appFlag := flag.String("flag", "foo", "pass in configuration")
	flag.Parse()

	fmt.Println("appFlag:\t", *appFlag)
	fmt.Println("Version:\t", Version)
	fmt.Println("CommitId\t", CommitId)
	fmt.Println("BuildTimestamp\t", BuildTimestamp)

	app := &wip.Wip{
		Name: *appFlag,
	}

	//run some code
	fmt.Println(app.One())

	fmt.Println(app.Two(3))

}
