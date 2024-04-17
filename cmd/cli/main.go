package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/mjdilworth/template/internal/api"
	"github.com/mjdilworth/template/internal/wip"
)

var Version = "development"
var CommitId string
var BuildTimestamp string

func main() {

	appFlag := flag.String("flag", "foo", "pass in configuration")

	serverPort := flag.String("port", ":8080", "specify the port the server listens on")

	flag.Parse()

	fmt.Println("appFlag:\t", *appFlag)
	fmt.Println("Version:\t", Version)
	fmt.Println("CommitId\t", CommitId)
	fmt.Println("BuildTimestamp\t", BuildTimestamp)

	app := &wip.Wip{
		Name: *appFlag,
	}

	fmt.Println(app.One())

	fmt.Println(app.Two(3))

	srv := &http.Server{
		Addr: *serverPort,
	}

	//simple
	http.HandleFunc("/health/", api.Health)
	http.HandleFunc("/", api.Root)
	http.HandleFunc("/secret/", api.Auth)
	http.HandleFunc("/spacepeeps/", api.Spacepeeps)

	th := api.TimeHandler(time.RFC1123)
	http.Handle("/time", th)

	http.HandleFunc("/help", api.Help)

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
