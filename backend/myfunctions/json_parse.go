package myfunctions

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
)

type web struct {
	Web WebStuff `json:"web"`
}

// IMPORTANT NOTE, VARIABLE NAME MUST START WITH A CAP :(((
type WebStuff struct {
	Client_id string `json:"client_id"`
	Project_id string `json:"project_id"`
	Auth_uri string `json:"auth_uri"`
	Token_uri string `json:"token_uri"`
	Auth_provider_x509_cert_url string `json:"clieauth_provider_x509_cert_urlnt_id"`
	Client_secret string `json:"client_secret"`
	Redirect_uris []string `json:"redirect_uris"`
	Javascript_origins []string `json:"javascript_origins"`
}

/*
	Client_id
	Project_id
	Auth_uri
	Token_uri
	Auth_provider_x509_cert_url
	Client_secret
	Redirect_uris
	Javascript_origins
*/

func ParseSecretJSON(fileLocation, wantedParameter string) string {

	fmt.Println("Printing File " + fileLocation)
	data, err := ioutil.ReadFile(fileLocation)

	if err != nil {
		fmt.Println("Err 1: ")
		fmt.Println(err)
	}

	var myJSONStuff web

	err = json.Unmarshal(data, &myJSONStuff)

	if err != nil {
		fmt.Println("Err 2: ")
		fmt.Println(err)
	}

	switch {
	case wantedParameter == "Client_id":
		return myJSONStuff.Web.Client_id
	case wantedParameter == "Project_id":
		return myJSONStuff.Web.Project_id
	case wantedParameter == "Auth_uri":
		return myJSONStuff.Web.Auth_uri
	case wantedParameter == "Auth_provider_x509_cert_url":
		return myJSONStuff.Web.Auth_provider_x509_cert_url
	case wantedParameter == "Client_secret":
		return myJSONStuff.Web.Client_secret
	case wantedParameter == "Redirect_uris":
		return myJSONStuff.Web.Redirect_uris[0]
	case wantedParameter == "Javascript_origins":
		return myJSONStuff.Web.Javascript_origins[0]
	default:
		fmt.Println("Invalid Parameter")
		return "0"
	}

}
