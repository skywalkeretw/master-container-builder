package main

import (
	"fmt"
	"net/http"
	"reflect"
	"sync"

	"github.com/gin-gonic/gin"
)

func test(nm string) {
	fmt.Println(nm)
}
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
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
