package main

import (
	"log"
	"time"

	"net/http"

	"github.com/go-chi/jwtauth"
	"github.com/pkg/errors"

	"github.com/titpetric/factory/resputil"
)

type JWT struct {
	tokenClaim string
	tokenAuth  *jwtauth.JWTAuth
}

func (JWT) new() *JWT {
	jwt := &JWT{
		tokenClaim: "user_id",
		tokenAuth:  jwtauth.New("HS256", []byte("K8UeMDPyb9AwFkzS"), nil),
	}
	log.Println("DEBUG JWT:", jwt.Encode("1"))
	return jwt
}

func (jwt *JWT) Encode(id string) string {
	claims := jwtauth.Claims{}.
		Set(jwt.tokenClaim, id).
		SetExpiryIn(30 * time.Second).
		SetIssuedNow()
	_, tokenString, _ := jwt.tokenAuth.Encode(claims)
	return tokenString
}

func (jwt *JWT) Verifier() func(http.Handler) http.Handler {
	return jwtauth.Verifier(jwt.tokenAuth)
}

func (jwt *JWT) Decode(r *http.Request) string {
	val, _ := jwt.Authenticate(r)
	return val
}

func (jwt *JWT) Authenticate(r *http.Request) (string, error) {
	token, claims, err := jwtauth.FromContext(r.Context())
	if err != nil || token == nil {
		return "", errors.Wrap(err, "Empty or invalid JWT")
	}
	if !token.Valid {
		return "", errors.New("Invalid JWT")
	}
	return claims[jwt.tokenClaim].(string), nil
}

func (jwt *JWT) Authenticator() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, err := jwt.Authenticate(r)
			if err != nil {
				resputil.JSON(w, err)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
