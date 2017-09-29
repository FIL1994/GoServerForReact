package main

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"net/http"
)

///main function
///starts up a gin-gonic server that serves static files
func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Recovery()) //recovers from panics and writes a 500
	//use gin.Default() to use Recovery and Logger

	//tell the router where the html files are
	r.LoadHTMLGlob("static/*.html")

	r.Use(static.Serve("/", static.LocalFile("./static", true)))

	//serve the index file if the directory or file does not exist
	r.NoRoute(func(c *gin.Context){
		c.HTML(http.StatusOK, "index.html", nil)
	})

	//run on port 8080
	r.Run(":8080")
}
