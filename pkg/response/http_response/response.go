package http_response

import (
	"github.com/gin-gonic/gin"

	"github.com/Hank-Kuo/go-example/pkg/customError/httpError"
	"github.com/Hank-Kuo/go-example/pkg/logger"
)

type response struct {
	StatusCode int
	Body       *ResponseBody
}

type ResponseBody struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func OK(statusCode int, message string, data interface{}) *response {
	return &response{
		statusCode,
		&ResponseBody{Status: "success", Message: message, Data: data},
	}
}

func Fail(err error, logger logger.Logger) *response {
	parseErr := httpError.ParseError(err)
	if parseErr.Detail != nil {
		logger.Error(parseErr.Detail)
	}
	return &response{
		parseErr.GetStatus(),
		&ResponseBody{Status: "fail", Message: parseErr.GetMessage()},
	}
}

func (r *response) ToJSON(c *gin.Context) {
	if r.Body.Status == "fail" {
		c.AbortWithStatusJSON(r.StatusCode, r.Body)
	} else {
		c.JSON(r.StatusCode, r.Body)
	}
}
