package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type ResponseBuilder struct{}

func (rb *ResponseBuilder) Success(c echo.Context, data interface{}) error {
	response := Response{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    data,
	}
	return c.JSON(http.StatusOK, response)
}

func (rb *ResponseBuilder) InvalidRequest(c echo.Context) error {
	response := Response{
		Status:  http.StatusBadRequest,
		Message: "Bad Request",
	}
	return c.JSON(http.StatusBadRequest, response)
}

func (rb *ResponseBuilder) InternalServiceError(c echo.Context) error {
	response := Response{
		Status:  http.StatusInternalServerError,
		Message: "Internal Service Error",
	}
	return c.JSON(http.StatusInternalServerError, response)
}

// Create an instance of ResponseBuilder
var ResponseBuilderInstance = &ResponseBuilder{}
