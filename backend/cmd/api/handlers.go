package main

import (
	"log"
	"errors"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"io"

	"github.com/go-chi/chi/v5"
)

// HelloPayload represents the response payload for route /home
type HelloPayload struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Version string `json:"version"`
}
// @Title Home
// @Summary show status of server
// @Description this is a method to show the server alive
// @Produces json
// @Success 200 {object} HelloPayload
// @Failure 400 {string} string "bad request"
// @Router  / [get]
func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	var payload = HelloPayload{
		Status:  "active",
		Message: "Authenticated Web Services up and running",
		Version: "1.0.0",
	}

	_ = app.writeJSON(w, http.StatusOK, payload)
}

type LoginPayload struct {
	Username    string `json:"username"`
	Password string `json:"password"`
}
// @Title login
// @Summary authenticate user for login 
// @Description authenticate user for login the server and get JWT tokens with openid-connect by keycloak
// @Produces json
// @param Username body LoginPayload true "user name registered at keycloak"
// @param Password body LoginPayload true "Password registered at keycloak" 
// @Success 200 {object} TokenPair
// @Failure 400 {string} string "bad request"
// @Router /login [post]
func (app *application) login(w http.ResponseWriter, r *http.Request) {
	var requestPayload LoginPayload

	// read json payload
	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	//log.Printf(" requestPayload :%s", requestPayload)

	KLoginPayload := &KLoginPayload{
		clientID : app.keycloak.clientID,
		username : requestPayload.Username,
		password : requestPayload.Password,
		grantType : "password",
		clientSecret : app.keycloak.clientSecret,
	}

	// authenticate user with Keycloak's payload settings
	tokens, err := app.client.login(KLoginPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}
	
	// return valid tokens and token expire time
	app.writeJSON(w, http.StatusAccepted, tokens)
}

// @Title logout
// @Summary logout user 
// @Description logout user
// @Success 200 
// @Failure 400 {string} string "bad request"
// @Router /logout [post]
func (app *application) logout(w http.ResponseWriter, r *http.Request) {
	//http.SetCookie(w, app.auth.GetExpiredRefreshCookie())
	w.WriteHeader(http.StatusAccepted)
}

// Organization represents the response payload with organization's information
type Organization struct {
	Id   int `json:"id"`
	Name string `json:"name"`
}

// @Title organization_list
// @Summary get opendata.swiss's organization list
// @Description get opendata.swiss's organization list with JWT token
// @Produces json
// @Success 200 {array} Organization
// @Failure 400 {string} string "bad request"
// @Router  /swissdata/organizations [get]
// @Security Bearer
func (app *application) organization_list(w http.ResponseWriter, r *http.Request) {
	type Organization_List struct {
		Help string `json:"help"`
		Success bool `json:"success"`
		Result []string `json:"result"`
	}

	requestURL := "https://ckan.opendata.swiss/api/3/action/organization_list"

	// get organization list
	resp, err := http.Get(requestURL)
	if err != nil {
		log.Printf("error making http request: %s\n", err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}
	
	if resp.StatusCode != http.StatusOK {
		app.errorJSON(w, errors.New("fail to making http request"), http.StatusBadRequest)
		return
	} 
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
    if err != nil {
        app.errorJSON(w, err, http.StatusBadRequest)
		return
    }

	var response Organization_List
    if err = json.Unmarshal(body, &response); err != nil {
        app.errorJSON(w, err, http.StatusBadRequest)
		return
    }

    var organizations []Organization

	organizations = append(organizations, Organization{Id: 0, Name: "None"})
	for index, element := range response.Result {
		organizations = append(organizations, Organization{Id: index + 1, Name: element})
	}

	_ = app.writeJSON(w, http.StatusOK, organizations)
}

// SwissData represents the response payload with swissdata's information
type SwissData struct {
	Id   int `json:"id"`
	Owner_org string `json:"owner_org"`
	Maintainer string  `json:"maintainer"`
	Issued string  `json:"issued"`
	Maintainer_email string  `json:"maintainer_email"`
	Download_url  string `json:"download_url"`
}

// @Title data_list
// @Summary get opendata.swiss's data list 
// @Description get opendata.swiss's data list with JWT token by organization name
// @param organization_name path string true "organization name" 
// @Produces json
// @Success 200 {array} SwissData
// @Failure 400 {string} string "bad request"
// @Router  /swissdata/datalist/{organization_name} [get]
// @Security Bearer
func (app *application) data_list(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	if name == "" {
		app.errorJSON(w,  errors.New("No name param"), http.StatusBadRequest)
		return
	}
	log.Printf("url param: %s", name)

	type swissData_List_Result_Results_Resurces struct {
		Download_url string  `json:"download_url"`
		X map[string]interface{}
	}

	type swissData_List_Result_Results struct {
		Owner_org string  `json:"owner_org"`
		Maintainer string  `json:"maintainer"`
		Issued string  `json:"issued"`
		Maintainer_email string  `json:"maintainer_email"`
		Resources []swissData_List_Result_Results_Resurces `json:"resources"`
		X map[string]interface{}
	}

	type swissData_List_Result struct {
		Count int `json:"count"`
		Sort string `json:"sort"`
		Results []swissData_List_Result_Results `json:"results"`
		X map[string]interface{}
	}

	type swissData_List struct {
		Help string `json:"help"`
		Success bool `json:"success"`
		Result swissData_List_Result `json:"result"`
	}

	requestURL := "https://ckan.opendata.swiss/api/3/action/package_search?fq=organization:" + name

	// get data list
	resp, err := http.Get(requestURL)
	if err != nil {
		log.Printf("error making http request: %s\n", err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}
	
	if resp.StatusCode != http.StatusOK {
		app.errorJSON(w, errors.New("fail to making http request"), http.StatusBadRequest)
		return
	} 

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("client: could not read response body: %s\n", err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	var response swissData_List
	
    if err = json.Unmarshal(body, &response); err != nil {
        app.errorJSON(w, err, http.StatusBadRequest)
		return
    }

	// convert to response payload
	var datalist []SwissData
	count := 0
	if (response.Result.Count > 0) {
		for _, element := range response.Result.Results { 
			for _, ele := range element.Resources {
				if ele.Download_url != "" {
					datalist = append(datalist, SwissData{
						Id: count, 
						Owner_org: element.Owner_org, 
						Maintainer: element.Maintainer,
						Issued: element.Issued,
						Maintainer_email: element.Maintainer_email,
						Download_url:ele.Download_url,
					})
					count ++
				}
			}
		}

	}
	_ = app.writeJSON(w, http.StatusOK, datalist)
}


