package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "github.com/gin-gonic/contrib/static"
  "fmt"
  "myfunctions" // local functions
  "strings"
)

// global variable
var youtubeAPI_key = ""

// This function is to
func main() {
	youtubeAPI_key = myfunctions.Setup()
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

  // Start listening and serving requests at localhost:5000
  router.Run(":5000")
}