package handlers

// CreateInstrumentConstants creates instrument constants (i.e. timeseries)
// func CreateInstrumentConstants(db *sqlx.DB) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		tc := models.TimeseriesCollection{}
// 		if err := c.Bind(&tc); err != nil {
// 			return c.JSON(http.StatusBadRequest, err)
// 		}
// 		// Get action information from context
// 		a, err := models.NewAction(c)
// 		if err != nil {
// 			return c.JSON(http.StatusInternalServerError, err)
// 		}

// 		nn, err := models.CreateInstrumentConstants(db, a, tc.Items)
// 		if err != nil {
// 			return c.JSON(http.StatusInternalServerError, err)
// 		}

// 		return c.JSON(http.StatusCreated, nn)
// 	}
// }
