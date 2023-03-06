# Youtube-API-Playlist
My plan for this project is to develop a small applications which utilize the Youtube API V3. My primary goal in using the Youtube API V3 is to obtain a users' personalized playlist and have it displayed on the frontend. Each video in the playlist will be displayed as a cards, where each card contains video information (may not be possible). 

To display this Card, I will be using React and MUI. I chose React and MUI because I have a bit experience working with it when I was doing Capstone Project at York 2021-2022. I would like to further enchance my learning and be able to develope this project to put onto my resume.

For the backend, I will be using Golong, mostly because I never used Golong and would like to get a better feel for it. 

tl;dr, a simple Youtube application that display all the videos in your playlist using Golong, Gin and Reactjs.

### Notes
Just a side note, this is my first time using Golang and the plan is to create the apps on the "Go" (heh) while learning how to use Go here and there. The pro is that this is might be the quicker way to get experience in programming with Go. The downside is that I might be miss the nuance in Go, but that's for another time I suppose. Also please do not use this on some sketchy places as this project offers no protection in user information.

## Developing Tools

### Go Basic
This is a personal note for myself in how to do each usual bits in a programming language (e.g. variables, for-loop, if statement, while, blah blha) in Go. By the way, gobyexample.com is a 
god send. I find the website effective.
https://gobyexample.com/environment-variables

### Go API Calling
TBD

### SQL
I'm implmenting SQLite because I need to find a way to save certain data. Such as the name of the video and the channel that the video is published at. I chose SQL because it is a good skill to try to have some experience with as well as I need a way to save those. 

### Postman
A useful developing tools that can help you create API calls and see if it works.

## Gin

### Warning
This is not a representation of my practical skill, but I'm just figuring out how to do a project myself since the projects that I have done was in collaboration or was for school and isn't really a kind of work that I can say its a project that I have done.

### Links Used
Purely for documentation for myself should I ever forget where I got such information from.
|Purpose of the link| Link Used |
|-|-|
|Quick starting:|https://developers.google.com/youtube/v3/quickstart/go|
|Sample of Go:|https://github.com/googleapis/google-api-go-client/tree/main/examples|
|Getting API Key in Go:|https://levelup.gitconnected.com/an-effective-way-to-read-environment-variables-in-go-7454e6613ae5|
|Example codes of Go:|https://gobyexample.com/|
|Example Youtube API Sample in Go:|https://github.com/youtube/api-samples/tree/master/go|
|Calling API with Go|https://tutorialedge.net/golang/consuming-restful-api-with-go/|
|Getting Gin to serve React with has different routes| https://github.com/gin-gonic/contrib/issues/90#issuecomment-990237367 |
|View Dashboard (this is only for me)|https://console.cloud.google.com/apis/dashboard?project=youtube-api-playlist-349120|
|HTTPS locally in Gin|https://github.com/gin-gonic/gin/issues/530|
|Create Cert and Key in .pem for Gin|https://support.microfocus.com/kb/doc.php?id=7013103 |
|HTTP/REST Calls with Youtube API|https://developers.google.com/youtube/v3/guides/auth/server-side-web-apps | 
|URL Encoding| https://golang.cafe/blog/how-to-url-encode-string-in-golang-example.html |
|Golang and JSON Parsing| https://www.golinuxcloud.com/golang-parse-json/ |
|Getting Google TOken|

## Prerequeistes
|Requirement|
|-|
|1. Golang|
|2. npm (to run the frontend)|
|3. Gin|
|4. Secret Key from Youtube|
|5. Authorized Test Users|

# Setup
|Steps|
|-|
|1. Run the Backend in the directory /backend/index.go and use the command `go run index.js`|
|2. If frontend not built yet, then go to the directory /frontend/ use the command `npm run build`. The backend serves the build package of the frontend|
|3. Assuming localhost hasn't been changed. The default localhost is localhost:5000. Go there if you want to see any changes|

#






