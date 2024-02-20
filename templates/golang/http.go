package main

import (
	"fmt"
	"net/http"
	"os"
	"sync"

	"github.com/gin-gonic/gin"
)


func HandleHttp(wg *sync.WaitGroup) {
	defer wg.Done()
	r := gin.Default()
	r.POST("/", func(c *gin.Context) {
		var rb map[string]interface{}

		// Bind JSON body to a map
		if err := c.BindJSON(&rb); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
		}

		value := {{FUNCTION_CALL}}
		//call function
		if reflect.TypeOf(value).Kind() == reflect.Struct {
			c.JSON(http.StatusOK, value)
		} else {
			c.Data(http.StatusOK, "text/plain", []byte(fmt.Sprint(value)))
		}

	})

	r.GET("/openapi", func(c *gin.Context) {
		jsonData, err := os.ReadFile("/root/openapi.json")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read JSON file"})
			return
		}
		c.Header("Content-Type", "application/json")
		c.Data(http.StatusOK, "application/json", jsonData)
	})

	r.GET("/asyncapi", func(c *gin.Context) {
		jsonData, err := os.ReadFile("/root/asyncapi.json")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read JSON file"})
			return
		}
		c.Header("Content-Type", "application/json")
		c.Data(http.StatusOK, "application/json", jsonData)
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
