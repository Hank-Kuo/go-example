package http_response

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type MockLogger struct{}

func (m *MockLogger) InitLogger()                                  {}
func (m *MockLogger) DPanic(args ...interface{})                   {}
func (m *MockLogger) DPanicf(template string, args ...interface{}) {}
func (m *MockLogger) Debug(args ...interface{})                    {}
func (m *MockLogger) Debugf(template string, args ...interface{})  {}
func (m *MockLogger) Info(args ...interface{})                     {}
func (m *MockLogger) Infof(atemplate string, rgs ...interface{})   {}
func (m *MockLogger) Warn(args ...interface{})                     {}
func (m *MockLogger) Warnf(template string, args ...interface{})   {}
func (m *MockLogger) Fatal(args ...interface{})                    {}
func (m *MockLogger) Fatalf(template string, args ...interface{})  {}
func (m *MockLogger) Error(args ...interface{})                    {}
func (m *MockLogger) Errorf(template string, args ...interface{})  {}

func TestOK(t *testing.T) {

	t.Run("PKG.HttpResponse.OK", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		response := OK(http.StatusOK, "Test OK", nil)
		response.ToJSON(c)

		assert.Equal(t, http.StatusOK, w.Code)

		expectedBody := `{"status":"success","message":"Test OK"}`
		assert.Equal(t, expectedBody, w.Body.String())
	})

}

func TestFail(t *testing.T) {
	t.Run("PKG.HttpResponse.Fail", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		mockError := errors.New("Test error")

		response := Fail(mockError, &MockLogger{})
		response.ToJSON(c)

		assert.Equal(t, http.StatusInternalServerError, w.Code)

		expectedBody := `{"status":"fail","message":"Internal Server Error"}`
		assert.Equal(t, expectedBody, w.Body.String())
	})

}

func TestToJSON(t *testing.T) {
	t.Run("PKG.HttpResponse.ToJSON", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		response := &response{
			StatusCode: http.StatusOK,
			Body: &ResponseBody{
				Status:  "success",
				Message: "OK",
				Data:    map[string]string{"key": "value"},
			},
		}

		response.ToJSON(c)

		assert.Equal(t, http.StatusOK, w.Code)

		expectedBody := `{"status":"success","message":"OK","data":{"key":"value"}}`
		assert.Equal(t, expectedBody, w.Body.String())
	})

}
