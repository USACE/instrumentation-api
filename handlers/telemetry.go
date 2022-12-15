package handlers

import (
	"log"
	"net/http"

	"github.com/USACE/instrumentation-api/models"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

func CreateOrUpdateDataLoggerMeasurements(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Parse the API key from the header, make sure its hash is in the database
		// This could also be done in the middleware but might not be as condusive
		// to change if we move to another API not using echo
		var df models.DataLoggerFile
		if err := c.Bind(&df); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		// TODO: Parse into structs for updating timeseries
		// Put timeseries measurments
		// stored, err := models.UpdateTimeseriesMeasurements(db, mcc.Items, &tw)
		// if err != nil {
		// 	return c.JSON(http.StatusBadRequest, err)
		// }
		log.Printf("Datafile: %+v\n", df)
		return c.JSON(http.StatusOK, df)
	}
}
