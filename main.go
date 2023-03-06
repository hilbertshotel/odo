package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"odo/dep"
	"odo/handlers"
)

func main() {
	// Initialize Dependencies
	d, err := dep.Init()
	if err != nil {
		panic(err)
	}

	// Parse Arguments
	cmd, name, err := parseArgs()
	if err != nil {
		d.Log.Error(err)
		return
	}

	// Handle Command
	switch cmd {
	case "info":
		fmt.Println(INFO)
	case "new":
		new(name)
	case "build":
		build(name)
	case "play":
		play(d)
	}
}

// CREATE NEW PROJECT IN GAMES FOLDER
func new(name string) {
	fmt.Println(name)
}

// TRANSPILE PROJECT FROM NIM TO JAVASCRIPT
func build(name string) {
	fmt.Println(name)
}

// START GAME SERVER
func play(d *dep.Dependencies) {
	// Initialize Server
	server := http.Server{
		Addr:         d.Cfg.HostAddr,
		Handler:      handlers.Mux(d),
		ReadTimeout:  time.Second * time.Duration(d.Cfg.ReadTimeout),
		WriteTimeout: time.Second * time.Duration(d.Cfg.WriteTimeout),
	}

	// Start Server
	serverError := make(chan error, 1)

	go func() {
		d.Log.Ok("SERVICE START @", d.Cfg.HostAddr)
		serverError <- server.ListenAndServe()
	}()

	// Handle Shutdown
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	select {

	case <-shutdown:
		server.Shutdown(context.Background())
		d.Log.Ok("SERVICE STOP")

	case err := <-serverError:
		d.Log.Error(err)
	}
}
