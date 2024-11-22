package endpoints

import (
	"encoding/base64"
	"log"
	"net/http"
	"strings"
)

func Login(w http.ResponseWriter, r *http.Request) {
	authorization := r.Header.Get("Authorization")

}

func parseAuthorizationHeader(authorization string) {
	elements := strings.Split(authorization, " ")
	b64 := elements[len(elements)-1]

	decoded, err := base64.RawStdEncoding.DecodeString(b64)

	if err != nil {
		log.Err
	}

	decode
}
