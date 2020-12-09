package utils

import (
	"encoding/base64"
)

func Encode(buff []byte) string {
	return base64.StdEncoding.EncodeToString(buff)
}

func Decode(b64 string) ([]byte, error) {
	buff, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		return nil, err
	}

	return buff, nil
}
