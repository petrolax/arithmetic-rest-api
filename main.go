package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)



func main() {
	router := gin.Default()
	
	router.GET("/api/add", Add)
	router.GET("/api/sub", Sub)
	router.GET("/api/mul", Mul)
	router.GET("/api/div", Div)

	router.Run(":80")
}

func Add(c *gin.Context) {
	aString := c.Query("a")
	bString := c.Query("b")

	if aString == "" || bString == "" {
		c.JSON(http.StatusNotAcceptable, map[string]interface{} { 
			"Success":false,
			"ErrCode":"406",
			"Value":nil,
		})
		return
	}

	a, err := strconv.ParseFloat(aString, 64)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, map[string]interface{} { 
			"Success":false,
			"ErrCode":"422",
			"Value":nil,
		})
		return
	}

	b, err := strconv.ParseFloat(bString, 64)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, map[string]interface{} { 
			"Success":false,
			"ErrCode":"422",
			"Value":nil,
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{} { 
		"Success": true,
		"ErrCode":"",
		"Value":a+b,
	})
}

func Sub(c *gin.Context) {
	aString := c.Query("a")
	bString := c.Query("b")

	if aString == "" || bString == "" {
		c.JSON(http.StatusNotAcceptable, map[string]interface{} { 
			"Success":false,
			"ErrCode":"406",
			"Value":nil,
		})
		return
	}

	a, err := strconv.ParseFloat(aString, 64)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, map[string]interface{} { 
			"Success":false,
			"ErrCode":"422",
			"Value":nil,
		})
		return
	}

	b, err := strconv.ParseFloat(bString, 64)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, map[string]interface{} { 
			"Success":false,
			"ErrCode":"422",
			"Value":nil,
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{} { 
		"Success": true,
		"ErrCode":"",
		"Value":a-b,
	})
}

func Mul(c *gin.Context) {
	aString := c.Query("a")
	bString := c.Query("b")

	if aString == "" || bString == "" {
		c.JSON(http.StatusNotAcceptable, map[string]interface{} { 
			"Success":false,
			"ErrCode":"406",
			"Value":nil,
		})
		return
	}

	a, err := strconv.ParseFloat(aString, 64)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, map[string]interface{} { 
			"Success":false,
			"ErrCode":"422",
			"Value":nil,
		})
		return
	}

	b, err := strconv.ParseFloat(bString, 64)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, map[string]interface{} { 
			"Success":false,
			"ErrCode":"422",
			"Value":nil,
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{} { 
		"Success": true,
		"ErrCode":"",
		"Value":a*b,
	})
}

func Div(c *gin.Context) {
	aString := c.Query("a")
	bString := c.Query("b")

	if aString == "" || bString == "" {
		c.JSON(http.StatusNotAcceptable, map[string]interface{} { 
			"Success":false,
			"ErrCode":"406",
			"Value":nil,
		})
		return
	}

	a, err := strconv.ParseFloat(aString, 64)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, map[string]interface{} { 
			"Success":false,
			"ErrCode":"422",
			"Value":nil,
		})
		return
	}

	b, err := strconv.ParseFloat(bString, 64)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, map[string]interface{} { 
			"Success":false,
			"ErrCode":"422",
			"Value":nil,
		})
		return
	}

	if b == 0 {
		c.JSON(http.StatusUnprocessableEntity, map[string]interface{} { 
			"Success":false,
			"ErrCode":"422",
			"Value":nil,
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{} { 
		"Success": true,
		"ErrCode":"",
		"Value":a/b,
	})
}