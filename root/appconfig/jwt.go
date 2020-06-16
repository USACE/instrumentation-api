package appconfig

import (
	"crypto/rsa"
	"log"
	"os"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var jwtVerifyKey = `-----BEGIN PUBLIC KEY-----
MIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEAxR6GTZ51RITOF9qNh1JH
GHEEHFj4kDVw1P5zumqW239XIdvn81sAslQm4ka0/e89q6Ci8WqRoJeoway0Ys0T
w83LcoQBdH461gVgzig+v8PZ8XiIkBLrkqXh6mspiBmOIWXIP6O6gqqJtZXEUBLf
8pKd8lmZu+wkUxUD5OzZMzoZoCOAnkP1MLVIZ9igS86XVgtR339zBeMeKYr9h2Fe
5uRgp0QvDjUxqLcPB+33ZGh8h1yVSPNjHBatU/mV/zENhPzdh9oZN+OMagHb05SC
JN06gT9LZNgfAiyYlXvbkACysfHG1k+Tw0bK7eN0pKxrh88a1/r90S4QQbgvo2Bw
ZQp9AtqX1VgCjjsTHBUrdmdt7qH6XFdUUlMk6OcCLU0pi0uXSqyvH9h4CkuYUUCI
13r8Ed7OB270Xh90hE7fj3Rb3o51FPI2FVOgvPp0f4HxnQzDe5nNPw7C1k620nZD
V6p4KXdJYkNZ6EqRNS2SY6iOFgXT9PNjCZu0Dgt/UbvecLboLJISiZ/9gceC9JTJ
MjKLreaAUM4ayrStgx5C8Nev6PLO8BpMoYM2Lb4Kt1PuuQxaDskeB7PBV3p6wS6X
jULmKThQMqJWNFxtKO1ZZaBOaXg50H0X+28RZdlPk6qgiFyK6LcVw8ZEemxk/3bk
dtc8yA3y/USzK7j6eu1XfOECAwEAAQ==
-----END PUBLIC KEY-----`

func parsePublicKey(key string) *rsa.PublicKey {

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(key))

	if err != nil {
		log.Printf(err.Error())
	}

	return publicKey

}

func skipper(c echo.Context) bool {
	if c.Request().Method == "GET" {
		return true
	}
	value, exists := os.LookupEnv("JWT_DISABLED")
	if exists && strings.ToUpper(value) == "TRUE" {
		return true
	}
	return false
}

// JWTConfig is JWT authentication configuration for this app
var JWTConfig = middleware.JWTConfig{
	// Skipper defines a function to skip middleware.
	Skipper: skipper,
	// Signing key to validate token.
	// Required.
	SigningKey: parsePublicKey(jwtVerifyKey),

	// Signing method, used to check token signing method.
	// Optional. Default value HS256.
	SigningMethod: "RS512",

	// Context key to store user information from the token into context.
	// Optional. Default value "user".
	// ContextKey:

	// Claims are extendable claims data defining token content.
	// Optional. Default value jwt.MapClaims
	// Claims: jwt.MapClaims,

	// TokenLookup is a string in the form of "<source>:<name>" that is used
	// to extract token from the request.
	// Optional. Default value "header:Authorization".
	// Possible values:
	// - "header:<name>"
	// - "query:<name>"
	// - "cookie:<name>"
	// TokenLookup:

	// AuthScheme to be used in the Authorization header.
	// Optional. Default value "Bearer".
	// AuthScheme: "Bearer"
}
