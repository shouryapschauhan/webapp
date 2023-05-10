package main

import (
	"log"
	"net/http"

	"github.com/alexedwards/scs/v2"
)

type application struct {
	Session *scs.SessionManager
}

func main() {
	// set up an app config
	app := application{}

	//get application routes
	mux := app.routes()

	//get a session manager
	app.Session = getSession()

	//print out a message
	log.Println("Starting server on port 8080...")

	//start the server
	log.Fatal(http.ListenAndServe(":8080", mux))
}
