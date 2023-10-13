package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/USACE/instrumentation-api/api/internal/message"
	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/USACE/instrumentation-api/api/internal/util"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// ListDataloggers godoc
//
//	@Summary lists dataloggers for a project
//	@Tags datalogger
//	@Produce json
//	@Param project_id path string true "project uuid" Format(uuid)
//	@Success 200 {array} model.Datalogger
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /dataloggers [get]
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
//	@Param project_id path string true "project uuid" Format(uuid)
//	@Param instrument_id path string true "instrument uuid" Format(uuid)
//	@Param datalogger body model.Datalogger true "datalogger payload"
//	@Success 200 {array} model.DataloggerWithKey
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /datalogger [post]
func (h *ApiHandler) CreateDatalogger(c echo.Context) error {
	ctx := c.Request().Context()
	n := model.Datalogger{}
	if err := c.Bind(&n); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	p := c.Get("profile").(model.Profile)
	n.Creator = p.ID

	if n.Name == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "valid `name` field required")
	}

	slugsTaken, err := h.DataloggerService.ListDataloggerSlugs(ctx)

	// Generate unique slug
	slug, err := util.NextUniqueSlug(n.Name, slugsTaken)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, message.InternalServerError)
	}
	n.Slug = slug

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
//	@Success 200 {object} model.DataloggerWithKey
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /datalogger/{datalogger_id}/key [put]
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
	u.Updater, u.UpdateDate = &profile.ID, &t

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
//	@Success 200 {object} model.Datalogger
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /datalogger/{datalogger_id} [get]
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
//	@Success 200 {object} model.Datalogger
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /datalogger/{datalogger_id} [put]
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
	u.Updater, u.UpdateDate = &profile.ID, &t

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
//	@Success 200 {object} map[string]interface{}
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /datalogger/{datalogger_id} [delete]
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
	d.Updater, d.UpdateDate = &profile.ID, &t

	if err := h.DataloggerService.DeleteDatalogger(ctx, d); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"id": dlID})
}

// GetDataloggerPreview godoc
//
//	@Summary gets the most recent datalogger preview by by datalogger id
//	@Tags datalogger
//	@Produce json
//	@Param datalogger_id path string true "datalogger uuid" Format(uuid)
//	@Success 200 {object} model.DataloggerPreview
//	@Failure 400 {object} echo.HTTPError
//	@Failure 404 {object} echo.HTTPError
//	@Failure 500 {object} echo.HTTPError
//	@Router /datalogger/{datalogger_id}/preview [get]
func (h *ApiHandler) GetDataloggerPreview(c echo.Context) error {
	dlID, err := uuid.Parse(c.Param("datalogger_id"))
	if err != nil {
		return err
	}

	// Get preview from c.Request().Context()
	preview, err := h.DataloggerService.GetDataloggerPreview(c.Request().Context(), dlID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, preview)
}
