package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/chunyukuo88/workouts/internal/app"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "go backend server port")
	flag.Parse()

	app, err := app.NewApplication()
	if err != nil {
		panic(err)
	}

	app.Logger.Println("app is running now ...")

	http.HandleFunc("/health", HealthCheck)
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		IdleTimeout:  time.Minute,
		WriteTimeout: 30 * time.Second,
	}

	app.Logger.Printf("we are running on port %d", port)

	err = server.ListenAndServe()
	if err != nil {
		app.Logger.Println("oh crap")
		app.Logger.Fatal(err)
	}
}
