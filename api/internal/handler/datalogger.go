package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/USACE/instrumentation-api/api/internal/messages"
	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// ListDataloggers
func (h ApiHandler) ListDataloggers(c echo.Context) error {
	pID := c.QueryParam("project_id")
	if pID != "" {

		pID, err := uuid.Parse(pID)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, messages.MalformedID)
		}

		dls, err := h.DataloggerStore.ListProjectDataloggers(c.Request().Context(), pID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return echo.NewHTTPError(http.StatusOK, dls)
	}

	dls, err := h.DataloggerStore.ListAllDataloggers(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, dls)
}

// CreateDatalogger
func (h ApiHandler) CreateDatalogger(c echo.Context) error {
	n := model.Datalogger{}
	if err := c.Bind(&n); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	p := c.Get("profile").(*model.Profile)
	n.Creator = p.ID

	if n.Name == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "valid `name` field required")
	}
	// Generate unique slug
	slug, err := h.DataloggerStore.CreateUniqueSlugDatalogger(c.Request().Context(), n.Name)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, messages.InternalServerError)
	}
	n.Slug = slug

	model, err := h.DataloggerStore.GetDataloggerModelName(c.Request().Context(), n.ModelID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("data logger model id %s not found", n.ModelID))
	}

	// check if datalogger with model and sn already exists and is not deleted
	exists, err := h.DataloggerStore.GetDataloggerIsActive(c.Request().Context(), model, n.SN)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if exists {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			"active data logger model with this model and serial number already exist",
		)
	}

	dl, err := h.DataloggerStore.CreateDatalogger(c.Request().Context(), n)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, messages.InternalServerError)
	}

	return c.JSON(http.StatusCreated, dl)
}

// CycleDataloggerKey
func (h ApiHandler) CycleDataloggerKey(c echo.Context) error {
	dlID, err := uuid.Parse(c.Param("datalogger_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, messages.MalformedID)
	}

	u := model.Datalogger{ID: dlID}

	if err := h.DataloggerStore.VerifyDataloggerExists(c.Request().Context(), dlID); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	profile := c.Get("profile").(*model.Profile)
	t := time.Now()
	u.Updater, u.UpdateDate = &profile.ID, &t

	dl, err := h.DataloggerStore.CycleDataloggerKey(c.Request().Context(), u)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, messages.InternalServerError)
	}

	return c.JSON(http.StatusOK, dl)
}

// GetDatalogger
func (h ApiHandler) GetDatalogger(c echo.Context) error {
	dlID, err := uuid.Parse(c.Param("datalogger_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, messages.MalformedID)
	}
	dl, err := h.DataloggerStore.GetOneDatalogger(c.Request().Context(), dlID)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, messages.NotFound)
	}

	return c.JSON(http.StatusOK, dl)
}

// UpdateDatalogger
func (h ApiHandler) UpdateDatalogger(c echo.Context) error {
	dlID, err := uuid.Parse(c.Param("datalogger_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, messages.MalformedID)
	}

	u := model.Datalogger{ID: dlID}
	if err := c.Bind(&u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if dlID != u.ID {
		return echo.NewHTTPError(http.StatusBadRequest, messages.MatchRouteParam("`id`"))
	}

	if err := h.DataloggerStore.VerifyDataloggerExists(c.Request().Context(), dlID); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	profile := c.Get("profile").(*model.Profile)
	t := time.Now()
	u.Updater, u.UpdateDate = &profile.ID, &t

	dlUpdated, err := h.DataloggerStore.UpdateDatalogger(c.Request().Context(), u)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, dlUpdated)
}

// DeleteDatalogger
func (h ApiHandler) DeleteDatalogger(c echo.Context) error {
	dlID, err := uuid.Parse(c.Param("datalogger_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, messages.MalformedID)
	}

	if err := h.DataloggerStore.VerifyDataloggerExists(c.Request().Context(), dlID); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	d := model.Datalogger{ID: dlID}
	profile := c.Get("profile").(*model.Profile)
	t := time.Now()
	d.Updater, d.UpdateDate = &profile.ID, &t

	if err := h.DataloggerStore.DeleteDatalogger(c.Request().Context(), d); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"id": dlID})
}

func (h ApiHandler) GetDataloggerPreview(c echo.Context) error {
	dlID, err := uuid.Parse(c.Param("datalogger_id"))
	if err != nil {
		return err
	}

	// Get preview from c.Request().Context()
	preview, err := h.DataloggerStore.GetDataloggerPreview(c.Request().Context(), dlID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, preview)
}
