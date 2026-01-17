package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"ecom/util"
	"net/http"
	"strings"
)

func (m *Middlewares) AuthJwt(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// parse jwt
		// parse header and payload or claims
		// hmac-sha-256 algorithm -> hash hmac(header, payload, secret key)
		// parse signature part from the jwt
		// if the signature and hash is same => forward to create products
		// otherwise 401 status code with Unauthorized

		header := r.Header.Get("Authorization")
		if header == "" {
			util.SendError(w, http.StatusUnauthorized, "Unauthorized")
			return
		}

		headerArr := strings.Split(header, " ")
		if len(headerArr) != 2 {
			util.SendError(w, http.StatusUnauthorized, "Unauthorized")
			return
		}

		accessToken := headerArr[1]

		tokenParts := strings.Split(accessToken, ".")
		if len(tokenParts) != 3 {
			util.SendError(w, http.StatusUnauthorized, "Unauthorized")
			return
		}

		jwtHeader := tokenParts[0]
		jwtPayload := tokenParts[1]
		jwtSignature := tokenParts[2]

		message := jwtHeader + "." + jwtPayload

		// cnf := config.GetConfig()

		byteArrSecret := []byte(m.cnf.JwtSecretKey)
		byteArrMessage := []byte(message)

		h := hmac.New(sha256.New, byteArrSecret)
		h.Write(byteArrMessage)

		hash := h.Sum(nil)
		hashSignature := util.Base64UrlEncode(hash)

		if hashSignature != jwtSignature {
			util.SendError(w, http.StatusUnauthorized, "Unauthorized f**kd off")
			return
		}

		next.ServeHTTP(w, r)
	})
}
