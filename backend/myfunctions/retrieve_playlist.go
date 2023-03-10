// https://developers.google.com/youtube/v3/guides/auth/server-side-web-apps#callinganapi

package myfunctions

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"net/url"
	"strconv" // needed to convert boolean to string
)

// Information to send over to YoutubeAPI
// struct is specifically designed to retreived only the playlist that the user
// own and nothing else.

// type MyPlaylist struct {
// 	Part string `json:"part"`
// 	Mine bool `json:"mine"`
// 	Access_token string `json:"access_token"`
// }

// Response JSON that we get from YouTube API after shooting out our JSON or whatever method that I used.
type MyPlaylistResponse struct {
	Kind string `json:"kind"`
	Etag string `json:"etag"`
	NextPageToken string `json:"nextPageToken"`
	PrevPageToken string `json:"PrevPageToken"`
	PageInfo PageInfo `json:"pageInfo"`
	Items []string `json:"items"`
}

type playlistInformation struct {


}

type PageInfo struct {
	TotalResults int `json:"totalResults"`
	resultsPerPage int `json:"resultsPerPage"`
}

/* 
Requires Youtube Playlist API and user Token to make authorization
Will be making HTTP call instead of using whatever gizmo that Youtube provide to make a request.

Requires 
> Part 
> Bool
> Access Token

*/

func RetrievePlayList(part string, access_token string, mine bool) {
	fmt.Println("Sanity Check - RetrievePlaylist")

	params := url.Values{}
	params.Add("part", part)
	params.Add("access_token", access_token)
	params.Add("mine", strconv.FormatBool(mine))
	GET_URL := "https://www.googleapis.com/youtube/v3/playlists?" + params.Encode()

	resp, err := http.Get(GET_URL)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	
	fmt.Println(string(body))
}

