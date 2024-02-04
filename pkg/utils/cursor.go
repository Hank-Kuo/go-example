package utils

import (
	"encoding/base64"
	"fmt"
)

func EncodeCursor(name string, id int) string {
	data := []byte(fmt.Sprintf("{'%s': %d}", name, id))
	return base64.StdEncoding.EncodeToString(data)
}

func DecodeCursor(name string, encodeStr string) (int, error) {
	data, err := base64.StdEncoding.DecodeString(encodeStr)

	if err != nil {
		return 0, err
	}

	var id int
	fmt.Sscanf(string(data), "{'"+name+"': %d}", &id)
	if id == 0 {
		id = 1
	}
	return id, nil
}
