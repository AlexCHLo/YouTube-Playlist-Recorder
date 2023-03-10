// https://developers.google.com/youtube/v3/guides/auth/server-side-web-apps

package myfunctions

import (
	"fmt"
	"net/url"
	"net/http"
	"bytes"
	"io/ioutil"
	"encoding/json"
	"log"
)

/*
Set Authorization Parameters

Parameter
client_id: Client ID obtained in API Console Credential Page.

redirect_url: Determine where the APi server redirects after the user completes the authorization flow.
** MUST MATCH WITH ONE OF THE AUTHORIZED REDIRECT URIs FOR THE OAUTH 2.0. CAN BE SET UP IN API Console Crediential Page. **

response_type: Determine whether the Google Oauth 2.0 should returns an authorization code.
** SET THE PARAMETER VALUE TO 'code' FOR WEB SERVER APPLICATIONS **

scope: A space-delimited list of scopes that identify the resources that your application could access 
on the user's behalf. These values inform the consent screen that Google displays to the user.
** OUR ONLY SCOPE IS 'https://www.googleapis.com/auth/youtube.readonly' **

access_type: Indicates whether your application can refresh access tokens when the user is not present at the browser.
Valid parameter values are online, which is the default value, and offline.
** SET PARAMTER TO OFFINE **

*/

func CreateRequestURL (client_id, redirect_uri, response_type, scope string) string {
	fmt.Println("CreateRequestURL Sanity Check")

	params := url.Values{}
	params.Add("client_id", client_id)
	params.Add("redirect_uri", redirect_uri)
	params.Add("response_type", response_type)
	params.Add("scope", scope)
	myURL := "https://accounts.google.com/o/oauth2/v2/auth?" + params.Encode()

	return myURL
}

// Our POST Response
type OAuth2Token struct {
	Access_token string `json:"access_token"`
	Expires_in int `json:"expires_in"` 
	Token_type string `json:"token_type"`
	Scope string `json:"scope"`
	Refresh_token string `json:"refresh_token"`
}

// Our POST Request body
type OAuth2Exchange struct {
	Client_id string `json:"client_id"`
	Client_secret string `json:"client_secret"`
	Code string `json:"code"`
	Grant_type string `json:"grant_type"`
	Redirect_uri string `json:"redirect_uri"`
}

// Has to be POST. The POST can be any shape or form like a x-www-form-ur-encoded or JSON. In this case,
// I would like to play around with more JSON, so I would send a POST with json for the body
func ExchangeAccessCodeForToken(client_id, client_secret, code, grant_type, redirect_uri, request_url string) string {
	fmt.Println("ExchangeAccessCodeForToken - Sanity Check")

	posturl := request_url


	// Creating a Post Body!
	exchangingCode := OAuth2Exchange{
		Client_id: client_id,
		Client_secret: client_secret,
		Code: code,
		Grant_type: grant_type,
		Redirect_uri: redirect_uri,
	}

	body, _ := json.Marshal(exchangingCode)

	resp, err := http.Post(posturl, "application/json", bytes.NewBuffer(body))

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	//Check response code, if New user is created then read response.
	if resp.StatusCode == http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			//Failed to read response.
			panic(err)
		}

		//Convert bytes to String and print
		jsonStr := string(body)
		fmt.Println("Response: ", jsonStr)

		// Obtain Token
		var token OAuth2Token
		err = json.Unmarshal(body, &token)

		if err != nil {
			panic(err)
		}

		log.Println(token.Access_token)

		return token.Access_token

	} else {
		//The status is not Created. print the error.
		fmt.Println("Get failed with error: ", resp.Status)
		panic("No Code Found")
	}

	

}
