package handler

import "github.com/labstack/echo/v4"

type Router struct {
	e  *echo.Echo
	ag AccessGroups
}

type AccessGroups struct {
	public  *echo.Group
	private *echo.Group
	cacOnly *echo.Group
}
