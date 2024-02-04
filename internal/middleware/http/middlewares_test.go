package http

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestNewGlobalMiddlewares(t *testing.T) {
	t.Run("Middleware.GlobalMiddlewares", func(t *testing.T) {
		r := gin.New()

		NewGlobalMiddlewares(r)

		req, _ := http.NewRequest("GET", "/notfound", nil)

		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
	})

}
