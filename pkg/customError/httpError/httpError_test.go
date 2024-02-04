package httpError

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHttpError(t *testing.T) {

	t.Run("PKG.Error", func(t *testing.T) {
		err := Err{Status: http.StatusNotFound, Message: "Not Found"}
		assert.Equal(t, "Not Found", err.Error(), "Error() method should return the correct message")
	})
	t.Run("PKG.Error.GetMessage", func(t *testing.T) {
		err := Err{Status: http.StatusNotFound, Message: "Not Found"}
		assert.Equal(t, "Not Found", err.GetMessage(), "GetMessage() should return the correct message")

		err = Err{Status: http.StatusNotFound}
		assert.Equal(t, "Internal server error", err.GetMessage(), "GetMessage() should return the default message")
	})
	t.Run("PKG.Error.GetStatus", func(t *testing.T) {
		err := Err{Status: http.StatusNotFound}
		assert.Equal(t, http.StatusNotFound, err.GetStatus(), "GetStatus() should return the correct status code")

		err = Err{}
		assert.Equal(t, http.StatusInternalServerError, err.GetStatus(), "GetStatus() should return the default status code")
	})

}

func TestNewError(t *testing.T) {

	t.Run("PKG.NewError", func(t *testing.T) {
		err := NewError(http.StatusBadRequest, "Bad request", errors.New("some error"))

		assert.Equal(t, http.StatusBadRequest, err.GetStatus())
		assert.Equal(t, "Bad request", err.GetMessage())
		assert.Equal(t, errors.New("some error"), err.Detail)
	})

	t.Run("PKG.NewNotFoundError", func(t *testing.T) {
		err := NewNotFoundError(errors.New("not found"))

		assert.Equal(t, http.StatusNotFound, err.GetStatus())
		assert.Equal(t, "Not Found", err.GetMessage())
		assert.Equal(t, errors.New("not found"), err.Detail)
	})
	t.Run("PKG.NewInternalServerError", func(t *testing.T) {
		err := NewInternalServerError(errors.New("Internal server error"))

		assert.Equal(t, http.StatusInternalServerError, err.GetStatus())
		assert.Equal(t, "Internal Server Error", err.GetMessage())
		assert.Equal(t, errors.New("Internal server error"), err.Detail)
	})

	t.Run("PKG.NewBadRequestError", func(t *testing.T) {
		err := NewBadRequestError(errors.New("Bad request error"))

		assert.Equal(t, http.StatusBadRequest, err.GetStatus())
		assert.Equal(t, "Bad request", err.GetMessage())
		assert.Equal(t, errors.New("Bad request error"), err.Detail)
	})

	t.Run("PKG.NewUnauthorizedError", func(t *testing.T) {
		err := NewUnauthorizedError(errors.New("Unauthorized error"))

		assert.Equal(t, http.StatusUnauthorized, err.GetStatus())
		assert.Equal(t, "Unauthorized", err.GetMessage())
		assert.Equal(t, errors.New("Unauthorized error"), err.Detail)
	})

	t.Run("PKG.NewRequestTimeoutError", func(t *testing.T) {
		err := NewRequestTimeoutError(errors.New("request time out error"))

		assert.Equal(t, http.StatusRequestTimeout, err.GetStatus())
		assert.Equal(t, "Request Timeout", err.GetMessage())
		assert.Equal(t, errors.New("request time out error"), err.Detail)
	})

	t.Run("PKG.NewBadQueryParamsError", func(t *testing.T) {
		err := NewBadQueryParamsError(errors.New("bad query params request error"))

		assert.Equal(t, http.StatusBadRequest, err.GetStatus())
		assert.Equal(t, "Invalid query params", err.GetMessage())
		assert.Equal(t, errors.New("bad query params request error"), err.Detail)
	})
}
