package httpError

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/lib/pq"
	errorss "github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestUnwrap(t *testing.T) {
	t.Parallel()
	t.Run("PKG.httpError.Unwrapped ErrorIsNil", func(t *testing.T) {
		err := errors.New("original error")
		result := unwrap(err)

		assert.Equal(t, err, result, "unwrap should return the original error when there are no wrapped errors")
	})

	t.Run("PKG.httpError.Unwrapped ErrorNotNil", func(t *testing.T) {
		wrappedErr := errors.New("wrapped error")
		err := fmt.Errorf("original error: %w", wrappedErr)
		result := unwrap(err)

		assert.Equal(t, wrappedErr, result, "unwrap should return the unwrapped error")
	})

	t.Run("PKG.httpError.Unwrapped NestedWrappedErrors", func(t *testing.T) {
		innerErr := errors.New("inner error")
		wrappedErr := fmt.Errorf("wrapped error: %w", innerErr)
		err := fmt.Errorf("original error: %w", wrappedErr)

		result := unwrap(err)

		assert.Equal(t, innerErr, result, "unwrap should return the most inner error")
	})

	t.Run("PKG.httpError.Unwrapped NilError", func(t *testing.T) {
		var err error
		result := unwrap(err)

		assert.Nil(t, result, "unwrap should return nil for nil error")
	})
}

func TestParseDBErrors(t *testing.T) {
	t.Run("PKG.httpError.ParseDBErrors", func(t *testing.T) {
		duplicateConstraint := "unique_constraint"

		pqError := &pq.Error{
			Code:       pq.ErrorCode("23505"),
			Constraint: "unique_constraint",
		}
		err := errorss.Wrap(pqError, "database error")

		result := parseDBErrors(err)

		assert.Equal(t, http.StatusConflict, result.GetStatus())
		expectedErrorMessage := fmt.Sprintf("This Field %s is used", duplicateConstraint)
		assert.Equal(t, expectedErrorMessage, result.GetMessage())
	})

}

func TestParseError(t *testing.T) {
	t.Parallel()

	t.Run("PKG.httpError.ParseError DeadlineExceededError", func(t *testing.T) {
		err := context.DeadlineExceeded
		result := ParseError(err)
		assert.Equal(t, http.StatusRequestTimeout, result.GetStatus())
		assert.Equal(t, "Request Timeout", result.GetMessage())
	})

	t.Run("PKG.httpError.ParseError FieldValidationError", func(t *testing.T) {
		err := errors.New("Field validation error")
		result := ParseError(err)

		assert.Equal(t, http.StatusBadRequest, result.GetStatus())
	})

	t.Run("PKG.httpError.ParseError UnmarshalError", func(t *testing.T) {
		err := errors.New("unmarshal error")
		result := ParseError(err)

		assert.Equal(t, http.StatusBadRequest, result.GetStatus())
	})

	t.Run("PKG.httpError.ParseError SqlError", func(t *testing.T) {
		err := errors.New("sql: error")
		result := ParseError(err)

		assert.Equal(t, http.StatusInternalServerError, result.GetStatus())
	})

	t.Run("PKG.httpError.ParseError DBError", func(t *testing.T) {
		err := errors.New("pg: error")
		result := ParseError(err)

		assert.Equal(t, http.StatusInternalServerError, result.GetStatus())
	})

	t.Run("PKG.httpError.ParseError CustomError", func(t *testing.T) {
		customErr := &Err{
			Status:  400,
			Message: "Bad Request",
			Detail:  errors.New("custom error detail"),
		}
		result := ParseError(customErr)
		assert.Equal(t, customErr, result)
	})

	t.Run("PKG.httpError.ParseError DefaultError", func(t *testing.T) {
		err := errors.New("unknown error")
		result := ParseError(err)

		assert.Equal(t, http.StatusInternalServerError, result.GetStatus())
		assert.Equal(t, "Internal Server Error", result.GetMessage())
	})
}

func TestParseSqlErrors(t *testing.T) {
	t.Parallel()

	t.Run("PKG.httpError.ParseSqlErrors ErrNoRows", func(t *testing.T) {
		err := sql.ErrNoRows
		result := parseSqlErrors(err)

		assert.Equal(t, http.StatusNotFound, result.GetStatus())
		assert.Equal(t, "Not Found", result.GetMessage())
	})

	t.Run("PKG.httpError.ParseSqlErrors ErrConnDone", func(t *testing.T) {
		err := sql.ErrConnDone
		result := parseSqlErrors(err)

		assert.Equal(t, http.StatusInternalServerError, result.GetStatus())
		assert.Contains(t, result.GetMessage(), "db connection is done")
	})

	t.Run("PKG.httpError.ParseSqlErrors ErrTxDone", func(t *testing.T) {
		err := sql.ErrTxDone
		result := parseSqlErrors(err)

		assert.Equal(t, http.StatusInternalServerError, result.GetStatus())
		assert.Contains(t, result.GetMessage(), "transaction operation error")
	})

	t.Run("PKG.httpError.ParseSqlErrors DefaultError", func(t *testing.T) {
		err := errors.New("unknown error")
		result := parseSqlErrors(err)

		assert.Equal(t, http.StatusInternalServerError, result.GetStatus())
		assert.Equal(t, "Internal Server Error", result.GetMessage())
	})
}

func TestParseMarshalError(t *testing.T) {
	t.Parallel()

	t.Run("PKG.httpError.ParseMarshalError UnmarshalTypeError", func(t *testing.T) {
		err := &json.UnmarshalTypeError{
			Value:  "unexpected value",
			Type:   reflect.TypeOf(0),
			Field:  "exampleField",
			Offset: 42,
		}
		result := parseMarshalError(err)
		assert.Equal(t, http.StatusBadRequest, result.GetStatus())
		assert.Contains(t, result.GetMessage(), "The Field: exampleField should be int")
	})

	t.Run("PKG.httpError.ParseMarshalError DefaultError", func(t *testing.T) {

		err := errors.New("unknown error")
		result := parseMarshalError(err)

		assert.Equal(t, http.StatusBadRequest, result.GetStatus())
		assert.Equal(t, "Bad request", result.GetMessage())
	})
}

func TestParseValidatorError(t *testing.T) {

	t.Run("PKG.httpError.ParseValidatorError ValidationErrors", func(t *testing.T) {

		err := parseValidatorError(errors.New("Error:Field validation for 'Status' failed on the 'required' tag"))

		assert.Equal(t, http.StatusBadRequest, err.GetStatus())
		assert.Contains(t, err.GetMessage(), "The Field: ")
	})
}
