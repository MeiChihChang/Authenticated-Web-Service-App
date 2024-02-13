package main

import (
	"fmt"
	"log"
	"time"
	"net/http"
	"crypto/tls"
	"os"

	"github.com/joho/godotenv"
)

const port = 8000

type application struct {
	client* Client
	keycloak Keycloak
}

func main() {
	// 
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	// set application config
	var app application
	
	tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true},}
    app.client = &Client{httpClient: &http.Client{Timeout: 6000 * time.Second,Transport: tr,},}
	// read from command line
	/*flag.StringVar(&app.keycloak.clientID, "clientID", "rest-golang-auth", "signing id")
	flag.StringVar(&app.keycloak.clientSecret, "clientSecret", "Nqe01UzPudKzXKbdqnLZ09XSLV3Qw0qw", "signing secret")
	flag.StringVar(&app.keycloak.userRole, "UserRole", "swiss-data-user", "signing issuer")
	flag.Parse()*/

	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal(err)
	}

	app.keycloak.clientID = os.Getenv("ClientID")
	app.keycloak.clientSecret = os.Getenv("ClientSecret")
	app.keycloak.userRole = os.Getenv("UserRole")
	

	log.Println("Starting application on port", port)

	// start a web server
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}
}
