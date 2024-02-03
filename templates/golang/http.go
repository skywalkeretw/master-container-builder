package main

import (
	"fmt"
	"net/http"
	"reflect"
	"sync"

	"github.com/gin-gonic/gin"
)

func HandleHttp(wg *sync.WaitGroup) {
	defer wg.Done()
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		value := {{FUNCTION_NAME}}
		//call function
		if reflect.TypeOf(value).Kind() == reflect.Struct {
			c.JSON(http.StatusOK, value)
		} else {
			c.Data(http.StatusOK, "text/plain", []byte(fmt.Sprint(value)))
		}

	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
