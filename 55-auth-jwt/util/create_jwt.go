package util

import (
	"encoding/base64"
	"encoding/json"
)

type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type Payload struct {
	Sub       string `json:"sub"`
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
	paloadB64 := base64UrlEncode(byteDataArr)

	byteArrSecret := []byte(secretKey)
	byteArrMessage := headerB64 + "." + paloadB64

}

func base64UrlEncode(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}
