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

func (rb *ResponseBuilder) Success(c echo.Context, data interface{}, message string) error {
	respMessage := "Success"
	if message != "" {
		respMessage = message
	}
	response := Response{
		Status:  http.StatusOK,
		Message: respMessage,
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

var ResponseBuilderInstance = &ResponseBuilder{}
