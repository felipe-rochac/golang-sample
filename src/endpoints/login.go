package endpoints

import (
	"authentication/constants"
	"encoding/base64"
	"errors"
	"net/http"
	"strings"
)

func Login(w http.ResponseWriter, r *http.Request) (*User, error) {
	authorization := r.Header.Get("Authorization")
	credentials, err := parseAuthorizationHeader(authorization)

	if err != nil {
		http.Error(w, er., constants.BadRequest)
	}
}

func parseAuthorizationHeader(authorization string) (*Credentials, error) {
	elements := strings.Split(authorization, " ")
	b64 := elements[len(elements)-1]

	if !strings.EqualFold(elements[0], "basic") {
		return nil, errors.New("Invalid authorization header")
	}

	decoded, err := base64.RawStdEncoding.DecodeString(b64)

	if err != nil {
		return nil, errors.New("Decoded string is invalid")
	}

	elements = strings.Split(string(decoded), ":")

	return &Credentials{
		Email:    elements[0],
		Password: elements[1],
	}, nil
}

type User struct {
	UserId  string
	Name    string
	Country string
	Enabled bool
}

type Credentials struct {
	Email    string
	Password string
}
