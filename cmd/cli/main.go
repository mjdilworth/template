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

	a := wip.New(*appFlag)

	fmt.Println(a.One())

	fmt.Println(a.Two(3))

}
