package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "github.com/gin-gonic/contrib/static"
  "google.golang.org/api/youtube/v3"
  "fmt"
  "myfunctions" // local functions
  "strings"
  "os"
)

// global variable
var youtubeAPI_key = ""
var tls = true // put false if you want to do http instead of https

// This function is to
func main() {
   // Get Youtube API Key
  youtubeAPI_key = myfunctions.Setup()

 // should we fail to get youtube api key, then stop the whole process.
  if (youtubeAPI_key == "") {
    os.Exit(3)    
  }
	fmt.Println(youtubeAPI_key)

  // Creates default gin router with Logger and Recovery middleware already attached
  router := gin.Default()

  /*
  Serve frontend static files
  note don't forget to add `..` at the beginning of the directory because we need to go back a file in order to
  access the frontend build.
  This is viewable in the / directory of the URL. React will take care of the front end.
  */
  router.Use(static.Serve("/", static.LocalFile("../frontend/build", true)))
  router.Use(static.Serve("/login", static.LocalFile("../frontend/build", true)))
  router.Use(static.Serve("/overview", static.LocalFile("../frontend/build", true)))

  /*
  Below is for specifically calling for API. Not serving frontend at all!
  Create API route group
  */

  api := router.Group("/api")
  {
	// Add /hello GET route to router and define route handler function
	// /api/hello
    api.GET("/hello", func(ctx *gin.Context) {
      ctx.JSON(http.StatusOK, gin.H {
		  "message" : "pong",
	  })
    })

    // Add /piss GET route for testing. please don't use this.
    api.GET("/piss", func(ctx *gin.Context) {
      ctx.JSON(http.StatusOK, gin.H {
		  "message" : "piss",
	  })
    })

    api.GET("/getClient", func(ctx *gin.Context) {
      client := myfunctions.GetClient(youtube.YoutubeReadonlyScope)
      fmt.Println(client)
    })

    api.GET("/getInformation", func(ctx *gin.Context) {
      client := myfunctions.GetClient(youtube.YoutubeReadonlyScope)
      state := ctx.Query("state")
      code := ctx.Query("code")
      ctx.Query("scope")
      fmt.Println()
    })
  }

    // Need to fix the 404 page not found
	router.NoRoute(func(c *gin.Context) {
		if !strings.HasPrefix(c.Request.RequestURI, "/api") {
      c.JSON(http.StatusOK, gin.H {
        "wrong place" : "buckeroo",
      })

		}
		//default 404 page not found
	})

  // router.NoRoute(func(ctx *gin.Context) { ctx.JSON(http.StatusNotFound, gin.H{}) })

  if tls {
    // Development SSL to ensure that I can get an HTTPS
    router.RunTLS(":5000", "./certifications/cert.pem", "./certifications/key.pem")
  } else {
    // Start listening and serving requests at localhost:5000
    router.Run(":5000")
  }

}