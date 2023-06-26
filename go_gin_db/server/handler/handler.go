package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/robert-min/performance_go_python_api/go_api_db/server/lib"
)

func getMovieHandler(c *gin.Context) {
	packages, err := lib.MysqlConnection()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": err.Error()})
	}

	data, _ := json.Marshal(packages)
	c.JSON(http.StatusOK, data)
}
