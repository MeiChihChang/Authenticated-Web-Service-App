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

// @title Authenticated Web Services
// @version 1.0
// @description This is a simple backend server for authenticated web services 
// @termsOfService http://localhost:8000
// @contact.name API Support
// @contact.email meichang@arch.nctu.edu.tw
// @host localhost:8000
// @BasePath /
// @securitydefinitions.apikey Bearer
// @in header
// @name Authorization
func main() {
	// 
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	// set application config
	var app application
	
	tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true},}
    app.client = &Client{httpClient: &http.Client{Timeout: 6000 * time.Second,Transport: tr,},}
	
    // read environment variable
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
