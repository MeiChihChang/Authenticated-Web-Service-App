package main

import (
	"log"
	"strings"
	"errors"
	"encoding/json"
	"context"
	"net/http"
	"net/url"

	oidc "github.com/coreos/go-oidc"
)

type keycloakService interface {
	login(*KLoginPayload) error
	
}

type Keycloak struct {
	clientID string
	clientSecret string
	userRole string
}

type KLoginPayload struct {
	clientID string
	username string
	password string
	grantType string
	clientSecret string
}

type Client struct {
	httpClient *http.Client
}

type TokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int `json:"expires_in"`
	RefreshExpiresIn    int `json:"refresh_expires_in"`
}

// authenticate user with openid-connect from keycloak
func (c *Client) login(kLoginPayload *KLoginPayload) (*TokenPair, error) {
	body := url.Values{
		"client_id": {kLoginPayload.clientID},
		"client_secret": {kLoginPayload.clientSecret},
		"grant_type": {kLoginPayload.grantType},
		"username": {kLoginPayload.username},
		"password": {kLoginPayload.password},
	}
	
	//log.Printf("body %s", body)
	encodedbody:= body.Encode()

	requestURL := "http://localhost:8080/realms/rest-golang/protocol/openid-connect/token"

    // get access token by openid-connect from keycloak	
	req, err := http.NewRequest(http.MethodPost, requestURL, strings.NewReader(encodedbody))
	if err != nil {
			log.Printf("client: could not create request: %s\n", err)
			return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		log.Printf("client: error making http request: %s\n", err)
		return nil, err
    }

	defer resp.Body.Close()
	
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("fail to keyloak login")
	} 

	// return access token & refresh token & expire time
	tokenpairs := &TokenPair{}
	json.NewDecoder(resp.Body).Decode(tokenpairs)
	//log.Printf("TokenPair %s", tokenpairs.ExpiresIn)	

	return tokenpairs, nil
}

type Res401Struct struct {
	Status   string `json:"status" example:"FAILED"`
	HTTPCode int    `json:"httpCode" example:"401"`
	Message  string `json:"message" example:"authorisation failed"`
}

//claims component of jwt contains mainy fields , we need only roles 
type Claims struct {
	ResourceAccess client `json:"resource_access,omitempty"`
	JTI            string `json:"jti,omitempty"`
}

type client struct {
	Restgolangauth clientRoles `json:"rest-golang-auth,omitempty"`
}

type clientRoles struct {
	Roles []string `json:"roles,omitempty"`
}


var RealmConfigURL string = "http://localhost:8080/realms/rest-golang"

// IsAuthorizedJWT verify token by online keycloak verifier and access permission with user role  
func (c *Client) IsAuthorizedJWT(w http.ResponseWriter, r *http.Request, clientID string, role string) error {
		// get auth header
		authHeader := r.Header.Get("Authorization")

		// sanity check
		if authHeader == "" {
			c.authorisationFailed("no auth header", w, r)
			return errors.New("no auth header")
		}

		// split the header on spaces
		headerParts := strings.Split(authHeader, " ")
		if len(headerParts) != 2 {
			c.authorisationFailed("invalid auth header", w, r)
			return errors.New("invalid auth header")
		}

		// check to see if we have the word Bearer
		if headerParts[0] != "Bearer" {
			c.authorisationFailed("invalid auth header", w, r)
			return errors.New("invalid auth header")
		}

		rawAccessToken := headerParts[1]

		ctx := oidc.ClientContext(context.Background(), c.httpClient)
		provider, err := oidc.NewProvider(ctx, RealmConfigURL)
		if err != nil {
			log.Printf("provider error :%v",err)
			c.authorisationFailed("authorisation failed while getting the provider: "+err.Error(), w, r)
			return err
		}

		// verify token
		oidcConfig := &oidc.Config{
			SkipClientIDCheck: true,
		}
		verifier := provider.Verifier(oidcConfig)
		idToken, err := verifier.Verify(ctx, rawAccessToken)
		if err != nil {
			log.Printf("verifier error :%v",err)
			c.authorisationFailed("authorisation failed while verifying the token: "+err.Error(), w, r)
			return err
		}

		// get claims to check role permission
		var IDTokenClaims Claims // ID Token payload is just JSON.
		if err := idToken.Claims(&IDTokenClaims); err != nil {
			log.Printf("Claims error :%v",err)
			c.authorisationFailed("claims : "+err.Error(), w, r)
			return err
		}
		log.Println(IDTokenClaims)
		//checking the roles
		user_access_roles := IDTokenClaims.ResourceAccess.Restgolangauth.Roles
		for _, b := range user_access_roles {
			if b == role {
				return nil
			}
		}

		log.Printf("role error :%v",user_access_roles)

		c.authorisationFailed("user not allowed to access this api", w, r)
		return errors.New("user not allowed to access this api")
}

// authorisationFailed writes response error code 
func (c *Client) authorisationFailed(message string, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusUnauthorized)
	data := Res401Struct{
		Status:   "FAILED",
		HTTPCode: http.StatusUnauthorized,
		Message:  message,
	}
	res, _ := json.Marshal(data)
	w.Write(res)
}