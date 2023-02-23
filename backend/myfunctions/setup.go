package myfunctions

import (
	"os"
	"fmt"
)

/*
setup() return youtube api key that is stored into your enviromental variable.
functions must start with a capital letter! otherwise golang would not be able to call the function outside of the
starting main() file.

Need to catch for error should the user did not get the youtube API Key
*/
func Setup() string {
	// Get API Key
	youtubeAPI_key := os.Getenv("YOUTUBE_V3_API")
	fmt.Printf("Youtube API Key = " + youtubeAPI_key + "\n")
	return youtubeAPI_key
}