package http

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCorsMiddleware(t *testing.T) {
	t.Run("Middleware.Cors", func(t *testing.T) {
		r := gin.New()
		r.Use(corsMiddleware())

		req, _ := http.NewRequest("GET", "/", nil)

		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, "*", w.Header().Get("Access-Control-Allow-Origin"))
		assert.Equal(t, "86400", w.Header().Get("Access-Control-Max-Age"))
		assert.Equal(t, "PUT, GET, POST, DELETE, OPTIONS", w.Header().Get("Access-Control-Allow-Methods"))
		assert.Equal(t, "*", w.Header().Get("Access-Control-Allow-Headers"))
		assert.Equal(t, "true", w.Header().Get("Access-Control-Allow-Credentials"))
	})
}
