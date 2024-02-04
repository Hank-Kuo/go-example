package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncodeCursor(t *testing.T) {
	t.Run("PKG.utils cursor encode", func(t *testing.T) {
		name := "example"
		id := 123
		expectedEncodedStr := "eydleGFtcGxlJzogMTIzfQ=="

		encodedStr := EncodeCursor(name, id)

		assert.Equal(t, expectedEncodedStr, encodedStr, "Encoded string does not match expected")
	})
	t.Run("PKG.utils cursor decode", func(t *testing.T) {
		name := "example"
		encodedStr := "eydleGFtcGxlJzogMTIzfQ==" // Replace with a valid encoded string for the given name and id

		expectedID := 123

		id, err := DecodeCursor(name, encodedStr)

		assert.NoError(t, err, "Unexpected error decoding cursor")
		assert.Equal(t, expectedID, id, "Decoded ID does not match expected")
	})

}
