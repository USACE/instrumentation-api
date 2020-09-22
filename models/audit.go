package models

import (
	"time"

	"github.com/labstack/echo/v4"
)

// AuditInfo holds common information about object creator and updater
type AuditInfo struct {
	Creator    int       `json:"creator"`
	CreateDate time.Time `json:"create_date" db:"create_date"`
	Updater    int       `json:"updater"`
	UpdateDate time.Time `json:"update_date" db:"update_date"`
}

// Action captures an API action, including user, user's roles,
// type (GET, POST, etc.), and Time
type Action struct {
	Actor int       `json:"actor"`
	Type  string    `json:"action"`
	Time  time.Time `json:"action_time"`
}

// NewAction returns a new action, taking values from echo.Context
func NewAction(c echo.Context) (*Action, error) {
	a := new(Action)
	a.Actor = c.Get("actor").(int)
	a.Type = c.Request().Method
	a.Time = time.Now()
	return a, nil
}
