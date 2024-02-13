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

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	var payload = struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Version string `json:"version"`
	}{
		Status:  "active",
		Message: "SwissData up and running",
		Version: "1.0.0",
	}

	_ = app.writeJSON(w, http.StatusOK, payload)
}

func (app *application) login(w http.ResponseWriter, r *http.Request) {
	// read json payload
	var requestPayload struct {
		Username    string `json:"username"`
		Password string `json:"password"`
	}

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

	//log.Printf(" KLoginPayload :%s", KLoginPayload)
	tokens, err := app.client.login(KLoginPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}
	
	app.writeJSON(w, http.StatusAccepted, tokens)
}

func (app *application) logout(w http.ResponseWriter, r *http.Request) {
	//http.SetCookie(w, app.auth.GetExpiredRefreshCookie())
	w.WriteHeader(http.StatusAccepted)
}

func (app *application) organization_list(w http.ResponseWriter, r *http.Request) {
	type Organization struct {
		Id   int `json:"id"`
		Name string `json:"name"`
	}
	type Organization_List struct {
		Help string `json:"help"`
		Success bool `json:"success"`
		Result []string `json:"result"`
	}

	requestURL := "https://ckan.opendata.swiss/api/3/action/organization_list"

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
	//log.Printf("response : %s", response)

    var organizations []Organization

	organizations = append(organizations, Organization{Id: 0, Name: "None"})
	for index, element := range response.Result {
		organizations = append(organizations, Organization{Id: index + 1, Name: element})
	}

	_ = app.writeJSON(w, http.StatusOK, organizations)
}

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

	type swissData struct {
		Id   int `json:"id"`
		Owner_org string `json:"owner_org"`
		Maintainer string  `json:"maintainer"`
		Issued string  `json:"issued"`
		Maintainer_email string  `json:"maintainer_email"`
		Download_url  string `json:"download_url"`
	}

	requestURL := "https://ckan.opendata.swiss/api/3/action/package_search?fq=organization:" + name

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
	//var p2 interface{} 
    if err = json.Unmarshal(body, &response); err != nil {
        app.errorJSON(w, err, http.StatusBadRequest)
		return
    }
	//log.Printf("response : %s", response)

	var datalist []swissData
	count := 0
	if (response.Result.Count > 0) {
		for _, element := range response.Result.Results { 
			for _, ele := range element.Resources {
				if ele.Download_url != "" {
					datalist = append(datalist, swissData{
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


