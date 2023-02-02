package middleware

import (
	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

var publicKey = `-----BEGIN PUBLIC KEY-----
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

// JWT is Fully Configured JWT Middleware to Support CWBI-Auth
func JWT(isDisabled bool, skipIfKey bool) echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningMethod: "RS512",
		KeyFunc: func(t *jwt.Token) (interface{}, error) {
			return jwt.ParseRSAPublicKeyFromPEM([]byte(publicKey))
		},
		Skipper: func(c echo.Context) bool {
			return isDisabled || (skipIfKey && c.QueryParam("key") != "")
		},
	})
}

// JWTMock is JWT Middleware
func JWTMock(isDisabled bool, skipIfKey bool) echo.MiddlewareFunc {
	return echojwt.WithConfig(
		echojwt.Config{
			SigningKey: []byte("mock"),
			Skipper: func(c echo.Context) bool {
				return isDisabled || (skipIfKey && c.QueryParam("key") != "")
			},
		},
	)
}
