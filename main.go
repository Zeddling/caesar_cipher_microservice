package main

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Data string `json:"response"`
}

func main() {
	//	Set router
	router := gin.Default()

	//	Init routes
	router.POST("/decrypt", DecryptHandler)
	router.POST("/encrypt", EncryptHandler)

	router.Run("localhost:8080")
}

/*
Params
	r -> Text to be encrypted.
	shift -> amount to shift letters with
	Function operates on the caesar equation c = (x + n) mod m
	where:
		c -> index of encrypted letter in the given alphabet
		x -> index of the actual letter
		n -> shift number
		m -> length of given alphabet
	Assumptions:
		Function assumes that the english alphabet is the basis for
		the encryption
*/
func caesar(r rune, shift int) rune {
	s := int(r) + shift
	if s > 'z' {
		return rune(s - 26)
	} else if s < 'a' {
		return rune(s + 26)
	}
	return rune(s)
}

/*
Handler expects two url parameter:
	1. data -> the encrypted text
	2. shift -> a negative degree used to shift data.
*/
var DecryptHandler = func(c *gin.Context) {
	//	Get URL params
	data := c.Query("data")
	shift, err := strconv.ParseInt(c.Query("shift"), 10, 64)

	//	Assert that shift is not empty
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "Shift cannot be empty",
			},
		)
		return
	}

	//	Assert that data is not empty
	if len(data) == 0 {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "Data cannot be empty",
			},
		)
		return
	}

	decrypted := strings.Map(func(r rune) rune {
		return caesar(r, -int(shift))
	}, data)

	c.IndentedJSON(
		http.StatusOK,
		Response{
			Data: decrypted,
		},
	)
}

/*
Handler expects two url parameters:
	1. data -> text to be encrypted
	2. shift -> a positive degree to shift data with
*/
var EncryptHandler = func(c *gin.Context) {
	//	Get URL params
	data := c.Query("data")
	shift, err := strconv.ParseInt(c.Query("shift"), 10, 64)

	//	Assert that shift is not empty
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "Shift cannot be empty",
			},
		)
		return
	}

	//	Assert that data is not empty
	if len(data) == 0 {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "Data cannot be empty",
			},
		)
		return
	}

	encrypted := strings.Map(func(r rune) rune {
		return caesar(r, int(shift))
	}, data)

	c.JSON(
		http.StatusOK,
		Response{
			Data: encrypted,
		},
	)
}
