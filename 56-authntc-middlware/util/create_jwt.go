package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
)

type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type Payload struct {
	Sub       int    `json:"sub"`
	FirstName string `json:"firstName"`
	Email     string `json:"email"`
}

func CreateJwt(data Payload, secretKey string) (string, error) {
	header := Header{
		Alg: "HS256",
		Typ: "JWT",
	}

	byteArr, err := json.Marshal(header)
	if err != nil {
		return "", err
	}
	headerB64 := base64UrlEncode(byteArr)

	byteDataArr, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	payloadB64 := base64UrlEncode(byteDataArr)

	byteArrSecret := []byte(secretKey)
	byteArrMessage := headerB64 + "." + payloadB64

	h := hmac.New(sha256.New, byteArrSecret)
	h.Write([]byte(byteArrMessage))

	signature := h.Sum(nil)
	signatureB64 := base64UrlEncode(signature)

	jwt := headerB64 + "." + payloadB64 + "." + signatureB64

	return jwt, nil
}

func base64UrlEncode(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}
