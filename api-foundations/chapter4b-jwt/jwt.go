package main

import (
	"log"

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

func (jwt *JWT) Encode(channel string) string {
	_, tokenString, _ := jwt.tokenAuth.Encode(jwtauth.Claims{jwt.tokenClaim: channel})
	return tokenString
}

func (jwt *JWT) Decode(r *http.Request) string {
	_, claims, _ := jwtauth.FromContext(r.Context())
	if val, ok := claims[jwt.tokenClaim]; ok {
		return val.(string)
	}
	return ""
}

func (jwt *JWT) Verifier() func(http.Handler) http.Handler {
	return jwtauth.Verifier(jwt.tokenAuth)
}

func (jwt *JWT) Authenticate(r *http.Request) (string, error) {
	token, claims, err := jwtauth.FromContext(r.Context())
	if err != nil || token == nil {
		return "", err
	}
	if !token.Valid {
		return "", errors.New("Invalid JWT")
	}
	return claims[jwt.tokenClaim].(string), nil
}

func (_ *JWT) Authenticator() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token, _, err := jwtauth.FromContext(r.Context())
			if err != nil {
				resputil.JSON(w, errors.Wrap(err, "Error validating JWT"))
				return
			}

			if token == nil || !token.Valid {
				resputil.JSON(w, errors.New("Empty or invalid JWT"))
				return
			}

			// Token is authenticated, pass it through
			next.ServeHTTP(w, r)
		})
	}
}
