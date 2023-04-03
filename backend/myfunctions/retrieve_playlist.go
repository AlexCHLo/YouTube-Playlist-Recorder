// https://developers.google.com/youtube/v3/guides/auth/server-side-web-apps#callinganapi

package myfunctions

import (
	"fmt"
	"log"
	"net/http"
	// "io/ioutil"
	"net/url"
	"strconv" // needed to convert boolean to string
	"encoding/json"
)

// Information to send over to YoutubeAPI
// struct is specifically designed to retreived only the playlist that the user
// own and nothing else.

/*
From what I see we have
{
	***
	***
	***
	items {


	}
}

*/

type Thumbnail struct {
	URL string `json:"url"`
}

type Thumbnails struct {
	Default  Thumbnail `json:"default"`
	Medium   Thumbnail `json:"medium"`
	High     Thumbnail `json:"high"`
	Standard Thumbnail `json:"standard"`
}

type Snippet struct {
	Title    string     `json:"title"`
	Thumbnails Thumbnails `json:"thumbnails"`
}

type Playlist struct {
	Snippet Snippet `json:"snippet"`
	Id      string  `json:"id"`
}

type PlaylistResponse struct {
	Items []Playlist `json:"items"`
}



/* 
Requires Youtube Playlist API and user Token to make authorization
Will be making HTTP call instead of using whatever gizmo that Youtube provide to make a request.

Requires 
> Part 
> Bool
> Access Token

*/

func RetrievePlayList(part string, access_token string, mine bool ) {
	fmt.Println("Sanity Check - RetrievePlaylist")

	params := url.Values{}
	params.Add("part", part)
	params.Add("access_token", access_token)
	params.Add("mine", strconv.FormatBool(mine))
	GET_URL := "https://www.googleapis.com/youtube/v3/playlists?" + params.Encode()
	fmt.Println(GET_URL)

	resp, err := http.Get(GET_URL)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	// parsing the response by set of playlist
	var playlistResponse PlaylistResponse
	err = json.NewDecoder(resp.Body).Decode(&playlistResponse)
	if err != nil {
		panic(err)
	}

	// Print the titles and the url of each images of each playlist
	for _, playlist := range playlistResponse.Items {
		fmt.Printf("Playlist Title: %s (ID: %s)\n", playlist.Snippet.Title, playlist.Id)
		fmt.Printf("Playlist Thumbnail URL: %s\n", playlist.Snippet.Thumbnails.Default.URL)
	}


	j, _ := json.Marshal(playlistResponse)
	log.Println(string(j))
  
	j, _ = json.MarshalIndent(playlistResponse, "", "  ")
	log.Println(string(j))


}

