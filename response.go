package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApiResponse struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type ApiError struct {
	Status int    `json:"status"`
	Code   string `json:"code"`
	Error  Error  `json:"error"`
}

/*
*

	Sends a formatted API response back to the client, containing both a `Status` and `Data` property on success,
	and a `Status`, `Code``

*
*/
func sendApiResponse(c *gin.Context, body interface{}) {
	statusCode := http.StatusOK

	c.IndentedJSON(statusCode, ApiResponse{
		Status: statusCode,
		Data:   body,
	})
}

/*
*

	Sends a formatted API error response back to the client, containing a `Status`, `Code`, and `Error` property

*
*/
func sendApiError(c *gin.Context, code string, error Error, status int) {
	statusCode := status | http.StatusInternalServerError
	c.IndentedJSON(statusCode, ApiError{
		Status: statusCode,
		Error:  error,
		Code:   code,
	})
}
