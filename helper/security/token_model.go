package security

import "github.com/golang-jwt/jwt/v5"

type TokenMyClaims struct {
	jwt.RegisteredClaims
	Id string `json:"id"`
}
