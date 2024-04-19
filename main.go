package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mjdilworth/template/internal/apploader"
	"github.com/mjdilworth/template/internal/server"
)

var Version = "development"
var CommitId string
var BuildTimestamp string

func main() {

	appFlag := flag.String("flag", "foo", "pass in configuration")
	daemon := flag.Bool("run as http daemon (true/false)", false)
	flag.Parse()

	fmt.Println("appFlag:\t", *appFlag)
	fmt.Println("Version:\t", Version)
	fmt.Println("CommitId\t", CommitId)
	fmt.Println("BuildTimestamp\t", BuildTimestamp)

	//create the application
	app := server.New()
	//load the applications
	al := apploader.New(app)

	//run the application
	if err := al.Run(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	//the application has ended
	al.LogMe("Service is ending", "key", 5)

}
