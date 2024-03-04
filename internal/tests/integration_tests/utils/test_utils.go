package utils

import (
	"encoding/json"
	"log"
)

func EncodePayload(payload any) []byte {

	payloadEncoded, err := json.Marshal(payload)
	if err != nil {
		log.Fatalf("Error to encode payload in create account controller test: %v", err)
	}

	return payloadEncoded
}
