package handler

import "github.com/gin-gonic/gin"

func Register() *gin.Engine {
	r := gin.Default()

	r.GET("/movies", getMovieHandler)
	return r
}
