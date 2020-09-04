package handlers

import (
	"api/root/models"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
)

// ListOpendcsSites returns all Instruments, represented as Opendcs Sites
func ListOpendcsSites(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		ss, err := models.ListOpendcsSites(db)
		if err != nil {
			return c.NoContent(http.StatusBadRequest)
		}
		return c.XMLPretty(http.StatusOK, ss, "  ")
	}
}
