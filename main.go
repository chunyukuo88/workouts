package main

import (
	"net/http"
	"time"

	"github.com/chunyukuo88/workouts/internal/app"
)

func main() {
	app, err := app.NewApplication()
	if err != nil {
		panic(err)
	}

	app.Logger.Println("app is running now ...")

	server := &http.Server{
		Addr:         "8080",
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	err = server.ListenAndServe()
	if err != nil {
		app.Logger.Println("oh crap")
		app.Logger.Fatal(err)
	}
}
