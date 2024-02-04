package http

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHealthzHandler(t *testing.T) {
	t.Run("Middleware.healthez", func(t *testing.T) {
		r := gin.New()
		Healthz(r)

		req, _ := http.NewRequest("GET", "/healthz", nil)

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
	})

}
