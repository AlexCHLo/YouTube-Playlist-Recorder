// https://developers.google.com/youtube/v3/guides/auth/server-side-web-apps

package myfunctions

import (
	"fmt"
	"net/url"
	"net/http"
	// "os"
	// "io/ioutil"
	// "log"
	"encoding/json"
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

type OAuth2Token struct {
	Access_token string `json:"access_token"`
	Expires_in string `json:"expires_in"`
	Token_type string `json:"token_type"`
	Scope string `json:"scope"`
	Refresh_token string `json:"refresh_token"`
}


// Has to be POST. The POST is a x-www-form-ur-encoded
func ExchangeAccessCodeForToken(client_id, client_secret, code, grant_type, redirect_uri, request_url string) {
	fmt.Println("ExchangeAccessCodeForToken - Sanity Check")

	params := url.Values{}
	
	params.Add("client_id", client_id)
	params.Add("client_secret", client_secret)
	params.Add("code", code)
	params.Add("grant_type", grant_type)
	params.Add("redirect_uri", redirect_uri)

	// encodedData := params.Encode()
	resp, err := http.PostForm(request_url, params)

	if err != nil {
		fmt.Println(err)
	}

	var res map[string]interface{}

    json.NewDecoder(resp.Body).Decode(&res)

    fmt.Println(res["form"])

}
