package utils

import (
	"encoding/json"
	"log"
)

func CompressToJSONByte(obj interface{}) []byte {
	raw, err := json.Marshal(obj)
	if err != nil {
		log.Fatal(err)
	}
	return raw
}
