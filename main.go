package main

import (
	"github.com/Anatol-e/grpcapp/greet/greet_client"
	"github.com/Anatol-e/grpcapp/greet/greet_server"
	log "github.com/sirupsen/logrus"
	"os"
)

const (
	server = "s"
	client = "c"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.TextFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)
}

func main() {
	appArgs := os.Args
	if len(appArgs) <= 1 {
		log.Fatal("Unknown application type! Must be server [s] or client [c]")
	}

	appType := appArgs[1]
	switch appType {
	case server:
		log.Info("Starting server...")
		greet_server.StartApplication()
	case client:
		log.Info("Starting client...")
		greet_client.StartApplication()
	}

}
