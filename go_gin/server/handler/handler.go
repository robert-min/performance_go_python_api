package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type pkgData struct {
	Name      string `json:"name"`
	MovieId   string `json:"movieId"`
	Rating    string `json:"rating"`
	TimeStamp string `json:"timestamp"`
}

// pathchData request to movie search server
func patchData(url string) ([]pkgData, error) {
	var packages []pkgData

	resp, err := http.Get(url)
	if err != nil {
		return packages, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return packages, err
	}

	err = json.Unmarshal(data, &packages)
	if err != nil {
		return packages, err
	}
	return packages, nil
}

func getMovieHandler(c *gin.Context) {
	url := "http://go_api_db:8080/search" // change server url
	packages, err := patchData(url)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": err.Error()})
	}

	jsonData, _ := json.Marshal(packages)

	c.JSON(http.StatusOK, jsonData)
}
