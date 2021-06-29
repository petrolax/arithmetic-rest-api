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

	router.Run(":8080")
}

func Add(c *gin.Context) {
	aString := c.Query("a")
	bString := c.Query("b")

	if aString == "" || bString == "" {
		c.JSON(http.StatusNotAcceptable, map[string]interface{}{
			"Success": false,
			"ErrCode": "One of the parameters haven't value or didn't exist",
			"Value":   nil,
		})
		return
	}

	a, err := strconv.ParseFloat(aString, 64)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"Success": false,
			"ErrCode": "Parameter 'a' is not a number",
			"Value":   nil,
		})
		return
	}

	b, err := strconv.ParseFloat(bString, 64)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"Success": false,
			"ErrCode": "Parameter 'b' is not a number",
			"Value":   nil,
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"Success": true,
		"ErrCode": "",
		"Value":   a + b,
	})
}

func Sub(c *gin.Context) {
	aString := c.Query("a")
	bString := c.Query("b")

	if aString == "" || bString == "" {
		c.JSON(http.StatusNotAcceptable, map[string]interface{}{
			"Success": false,
			"ErrCode": "One of the parameters haven't value or didn't exist",
			"Value":   nil,
		})
		return
	}

	a, err := strconv.ParseFloat(aString, 64)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"Success": false,
			"ErrCode": "Parameter 'a' is not a number",
			"Value":   nil,
		})
		return
	}

	b, err := strconv.ParseFloat(bString, 64)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"Success": false,
			"ErrCode": "Parameter 'b' is not a number",
			"Value":   nil,
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"Success": true,
		"ErrCode": "",
		"Value":   a - b,
	})
}

func Mul(c *gin.Context) {
	aString := c.Query("a")
	bString := c.Query("b")

	if aString == "" || bString == "" {
		c.JSON(http.StatusNotAcceptable, map[string]interface{}{
			"Success": false,
			"ErrCode": "One of the parameters haven't value or didn't exist",
			"Value":   nil,
		})
		return
	}

	a, err := strconv.ParseFloat(aString, 64)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"Success": false,
			"ErrCode": "Parameter 'a' is not a number",
			"Value":   nil,
		})
		return
	}

	b, err := strconv.ParseFloat(bString, 64)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"Success": false,
			"ErrCode": "Parameter 'b' is not a number",
			"Value":   nil,
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"Success": true,
		"ErrCode": "",
		"Value":   a * b,
	})
}

func Div(c *gin.Context) {
	aString := c.Query("a")
	bString := c.Query("b")

	if aString == "" || bString == "" {
		c.JSON(http.StatusNotAcceptable, map[string]interface{}{
			"Success": false,
			"ErrCode": "One of the parameters haven't value or didn't exist",
			"Value":   nil,
		})
		return
	}

	a, err := strconv.ParseFloat(aString, 64)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"Success": false,
			"ErrCode": "Parameter 'a' is not a number",
			"Value":   nil,
		})
		return
	}

	b, err := strconv.ParseFloat(bString, 64)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"Success": false,
			"ErrCode": "Parameter 'b' is not a number",
			"Value":   nil,
		})
		return
	}

	if b == 0 {
		c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"Success": false,
			"ErrCode": "You can't divide by zero",
			"Value":   nil,
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"Success": true,
		"ErrCode": "",
		"Value":   a / b,
	})
}
