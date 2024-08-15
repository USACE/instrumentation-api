package middleware

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func mwJWT(publicKey, signingMethod string, isDisabled, skipIfKey, mock bool) echo.MiddlewareFunc {
	pk := fmt.Sprintf("-----BEGIN PUBLIC KEY-----\n%s\n-----END PUBLIC KEY-----", publicKey)
	if mock {
		return echojwt.WithConfig(echojwt.Config{
			SigningKey: []byte("mock"),
			Skipper: func(c echo.Context) bool {
				return isDisabled || (skipIfKey && c.QueryParam("key") != "")
			},
		})
	}
	return echojwt.WithConfig(echojwt.Config{
		SigningMethod: signingMethod,
		KeyFunc: func(t *jwt.Token) (interface{}, error) {
			return jwt.ParseRSAPublicKeyFromPEM([]byte(pk))
		},
		Skipper: func(c echo.Context) bool {
			return isDisabled || (skipIfKey && c.QueryParam("key") != "")
		},
	})
}

func (m *mw) JWTSkipIfKey(next echo.HandlerFunc) echo.HandlerFunc {
	jwtmw := mwJWT(m.cfg.AuthPublicKey, m.cfg.AuthSigningMethod, m.cfg.AuthDisabled, true, m.cfg.AuthJWTMocked)
	return jwtmw(next)
}

func (m *mw) JWT(next echo.HandlerFunc) echo.HandlerFunc {
	jwtmw := mwJWT(m.cfg.AuthPublicKey, m.cfg.AuthSigningMethod, m.cfg.AuthDisabled, false, m.cfg.AuthJWTMocked)
	return jwtmw(next)
}
