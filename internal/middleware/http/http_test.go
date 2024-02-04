package http

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHTTPNotFoundHandler(t *testing.T) {
	t.Run("Middleware.NotFound", func(t *testing.T) {
		r := gin.New()

		r.GET("/notfound", httpNotFound)
		req, _ := http.NewRequest("GET", "/notfound", nil)

		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
	})

}
