package api

import (
	"net/http"
	"strings"
)

func GetTokenFromHeader(r http.Request) (string, bool) {
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		return "", false
	}

	return strings.TrimPrefix(tokenString, "Bearer "), true
}
