package handler

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/USACE/instrumentation-api/api/internal/message"
	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// ListDataloggers godoc
//
//	@Summary lists dataloggers for a project
//	@Tags datalogger
//	@Produce json
//	@Param key query string false "api key"
//	@Success 200 {array} model.Datalogger
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /dataloggers [get]
//	@Security Bearer
func (h *ApiHandler) ListDataloggers(c echo.Context) error {
	pID := c.QueryParam("project_id")
	if pID != "" {
		pID, err := uuid.Parse(pID)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, message.MalformedID)
		}

		dls, err := h.DataloggerService.ListProjectDataloggers(c.Request().Context(), pID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return echo.NewHTTPError(http.StatusOK, dls)
	}

	dls, err := h.DataloggerService.ListAllDataloggers(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, dls)
}

// CreateDatalogger godoc
//
//	@Summary creates a datalogger
//	@Tags datalogger
//	@Accept json
//	@Produce json
//	@Param datalogger body model.Datalogger true "datalogger payload"
//	@Param key query string false "api key"
//	@Success 200 {array} model.DataloggerWithKey
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /datalogger [post]
//	@Security Bearer
func (h *ApiHandler) CreateDatalogger(c echo.Context) error {
	ctx := c.Request().Context()
	n := model.Datalogger{}
	if err := c.Bind(&n); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	p := c.Get("profile").(model.Profile)
	n.CreatorID = p.ID

	if n.Name == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "valid `name` field required")
	}

	model, err := h.DataloggerService.GetDataloggerModelName(ctx, n.ModelID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("data logger model id %s not found", n.ModelID))
	}

	// check if datalogger with model and sn already exists and is not deleted
	exists, err := h.DataloggerService.GetDataloggerIsActive(ctx, model, n.SN)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if exists {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			"active data logger model with this model and serial number already exist",
		)
	}

	dl, err := h.DataloggerService.CreateDatalogger(ctx, n)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, message.InternalServerError)
	}

	return c.JSON(http.StatusCreated, dl)
}

// CycleDataloggerKey godoc
//
//	@Summary deletes and recreates a datalogger api key
//	@Tags datalogger
//	@Produce json
//	@Param datalogger_id path string true "datalogger uuid" Format(uuid)
//	@Param key query string false "api key"
//	@Success 200 {object} model.DataloggerWithKey
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /datalogger/{datalogger_id}/key [put]
//	@Security Bearer
func (h *ApiHandler) CycleDataloggerKey(c echo.Context) error {
	ctx := c.Request().Context()
	dlID, err := uuid.Parse(c.Param("datalogger_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, message.MalformedID)
	}

	u := model.Datalogger{ID: dlID}

	if err := h.DataloggerService.VerifyDataloggerExists(ctx, dlID); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	profile := c.Get("profile").(model.Profile)
	t := time.Now()
	u.UpdaterID, u.UpdateDate = &profile.ID, &t

	dl, err := h.DataloggerService.CycleDataloggerKey(ctx, u)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, message.InternalServerError)
	}

	return c.JSON(http.StatusOK, dl)
}

// GetDatalogger godoc
//
//	@Summary gets a datalogger by id
//	@Tags datalogger
//	@Produce json
//	@Param datalogger_id path string true "datalogger uuid" Format(uuid)
//	@Param key query string false "api key"
//	@Success 200 {object} model.Datalogger
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /datalogger/{datalogger_id} [get]
//	@Security Bearer
func (h *ApiHandler) GetDatalogger(c echo.Context) error {
	dlID, err := uuid.Parse(c.Param("datalogger_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, message.MalformedID)
	}
	dl, err := h.DataloggerService.GetOneDatalogger(c.Request().Context(), dlID)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, message.NotFound)
	}

	return c.JSON(http.StatusOK, dl)
}

// UpdateDatalogger godoc
//
//	@Summary updates a datalogger
//	@Tags datalogger
//	@Produce json
//	@Param datalogger_id path string true "datalogger uuid" Format(uuid)
//	@Param datalogger body model.Datalogger true "datalogger payload"
//	@Param key query string false "api key"
//	@Success 200 {object} model.Datalogger
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /datalogger/{datalogger_id} [put]
//	@Security Bearer
func (h *ApiHandler) UpdateDatalogger(c echo.Context) error {
	ctx := c.Request().Context()
	dlID, err := uuid.Parse(c.Param("datalogger_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, message.MalformedID)
	}

	u := model.Datalogger{ID: dlID}
	if err := c.Bind(&u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if dlID != u.ID {
		return echo.NewHTTPError(http.StatusBadRequest, message.MatchRouteParam("`id`"))
	}

	if err := h.DataloggerService.VerifyDataloggerExists(ctx, dlID); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	profile := c.Get("profile").(model.Profile)
	t := time.Now()
	u.UpdaterID, u.UpdateDate = &profile.ID, &t

	dlUpdated, err := h.DataloggerService.UpdateDatalogger(ctx, u)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, dlUpdated)
}

// DeleteDatalogger godoc
//
//	@Summary deletes a datalogger by id
//	@Tags datalogger
//	@Produce json
//	@Param datalogger_id path string true "datalogger uuid" Format(uuid)
//	@Param key query string false "api key"
//	@Success 200 {object} map[string]interface{}
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /datalogger/{datalogger_id} [delete]
//	@Security Bearer
func (h *ApiHandler) DeleteDatalogger(c echo.Context) error {
	ctx := c.Request().Context()
	dlID, err := uuid.Parse(c.Param("datalogger_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, message.MalformedID)
	}

	if err := h.DataloggerService.VerifyDataloggerExists(ctx, dlID); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	d := model.Datalogger{ID: dlID}
	profile := c.Get("profile").(model.Profile)
	t := time.Now()
	d.UpdaterID, d.UpdateDate = &profile.ID, &t

	if err := h.DataloggerService.DeleteDatalogger(ctx, d); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"id": dlID})
}

// GetDataloggerTablePreview godoc
//
//	@Summary gets the most recent datalogger preview by by datalogger id
//	@Tags datalogger
//	@Produce json
//	@Param datalogger_id path string true "datalogger uuid" Format(uuid)
//	@Param datalogger_table_id path string true "datalogger table uuid" Format(uuid)
//	@Param key query string false "api key"
//	@Success 200 {object} model.DataloggerTablePreview
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /datalogger/{datalogger_id}/tables/{datalogger_table_id}/preview [get]
//	@Security Bearer
func (h *ApiHandler) GetDataloggerTablePreview(c echo.Context) error {
	_, err := uuid.Parse(c.Param("datalogger_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, message.MissingQueryParameter("datalogger_id"))
	}
	dataloggerTableID, err := uuid.Parse(c.Param("datalogger_table_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, message.MissingQueryParameter("datalogger_table_id"))
	}
	preview, err := h.DataloggerService.GetDataloggerTablePreview(c.Request().Context(), dataloggerTableID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return echo.NewHTTPError(http.StatusNotFound, message.NotFound)
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, preview)
}

// ResetDataloggerTableName godoc
//
//	@Summary resets a datalogger table name to be renamed by incoming telemetry
//	@Tags datalogger
//	@Produce json
//	@Param datalogger_id path string true "datalogger uuid" Format(uuid)
//	@Param datalogger_table_id path string true "datalogger table uuid" Format(uuid)
//	@Param key query string false "api key"
//	@Success 200 {object} model.DataloggerTablePreview
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /datalogger/{datalogger_id}/tables/{datalogger_table_id}/name [put]
//	@Security Bearer
func (h *ApiHandler) ResetDataloggerTableName(c echo.Context) error {
	_, err := uuid.Parse(c.Param("datalogger_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, message.MissingQueryParameter("datalogger_id"))
	}
	dataloggerTableID, err := uuid.Parse(c.Param("datalogger_table_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, message.MissingQueryParameter("datalogger_table_id"))
	}
	if err := h.DataloggerService.ResetDataloggerTableName(c.Request().Context(), dataloggerTableID); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"datalogger_table_id": dataloggerTableID})
}
