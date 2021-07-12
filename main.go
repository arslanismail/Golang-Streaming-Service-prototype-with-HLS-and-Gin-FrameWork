package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var Movies = map[string]string{
	//  Add here the  movies list replace this with the movie list fetched from Database
	"tseries": "/tseries",
}

type Movie struct {
	Name string
}

const (
	mediaRoot     = "assets/media"
	basicM3U8Name = "index.m3u8"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	router.GET("/", index)
	router.GET("/stream/:movie/", stream)
	router.GET("/stream/:movie/:segName", stream)
	http.ListenAndServe(":8080", router)
}

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title":  "Streaming Service Golang",
		"Movies": Movies,
	})
}

func stream(c *gin.Context) {
	// Streaming Service
	movieName, _ := c.Params.Get("movie")
	// movie = Movie{movieName}
	fmt.Println(movieName)
	segName, _ := c.Params.Get("segName")
	fmt.Println(segName)
	if segName == "" {
		useM3U8(c.Writer, c.Request, movieName)
	} else {
		fmt.Println("segname is ", segName)
		useHlsTs(c.Writer, c.Request, segName, movieName)
	}
}

func useHlsTs(w http.ResponseWriter, r *http.Request, segname string, mId string) {
	mediaFile := fmt.Sprintf("%s/hls/%s", getMediaBase(mId), segname)
	http.ServeFile(w, r, mediaFile)
	w.Header().Set("Content-Type", "video/MP2T")
}
func useM3U8(w http.ResponseWriter, r *http.Request, mId string) {
	mediaFile := fmt.Sprintf("%s/hls/%s", getMediaBase(mId), basicM3U8Name)
	http.ServeFile(w, r, mediaFile)
	w.Header().Set("Content-Type", "application/x-mpegURL")
}
func getMediaBase(mId string) string {
	return fmt.Sprintf("%s/%s", mediaRoot, mId)
}
