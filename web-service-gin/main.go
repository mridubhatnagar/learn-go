package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)


func main() {
	router := gin.Default()
	router.GET("/", display)
	router.Run("localhost:8080")
}

func display(c *gin.Context) {
	var s = map[string]string{"msg": "Hello World"}
	c.IndentedJSON(http.StatusOK, s)
}