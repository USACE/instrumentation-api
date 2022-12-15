package handlers

import (
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

// Example payload from data logger
// {
// 	"head": {
// 	  "transaction": 0,
// 	  "signature": 1182,
// 	  "environment": {
// 		"station_name": "6239",
// 		"table_name": "Test",
// 		"model": "CR6",
// 		"serial_no": "6239",
// 		"os_version": "CR6.Std.12.01",
// 		"prog_name": "CPU:httpex.CR6"
// 	  },
// 	  "fields": [
// 		{
// 		  "name": "batt_volt_Min",
// 		  "type": "xsd:float",
// 		  "units": "Volts",
// 		  "process": "Min",
// 		  "settable": false
// 		},
// 		{
// 		  "name": "PanelT",
// 		  "type": "xsd:float",
// 		  "units": "Deg_C",
// 		  "process": "Smp",
// 		  "settable": false
// 		}
// 	  ]
// 	},
// 	"data": [
// 	  {
// 		"time": "2022-12-15T13:24:00",
// 		"no": 0,
// 		"vals": [
// 		  12.15,
// 		  22.71
// 		]
// 	  }
// 	]
// }

func CreateOrUpdateDataLoggerMeasurements(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Parse the API key from the header, make sure its hash is in the database
		// This could also be done in the middleware but might not be as condusive
		// to change if we move to another API not using echo

		// Get TOA5 Body from request
		log.Printf("request: %v", c)
		return c.JSON(http.StatusOK, c)
	}
}
