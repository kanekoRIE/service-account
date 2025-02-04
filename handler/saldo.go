package handler

import (
	"database/sql"
	"net/http"
	"service-account/helper"
	"service-account/logger"
	"service-account/model"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func Saldo(c echo.Context) error {
	var req struct {
		NoRekening string `param:"no_rekening" validate:"required,max=15"`
	}

	if err := c.Bind(&req); err != nil {
		logger.LogError("Failed to bind request", logrus.Fields{"error": err})
		return c.JSON(http.StatusBadRequest, map[string]string{"remark": "Invalid input"})
	}

	if err := helper.Validate.Struct(req); err != nil {
		var validationErrors []string
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, err.Field())
		}
		logger.LogError("Invalid Field: "+strings.Join(validationErrors, ", "), logrus.Fields{
			"error": err,
		})
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"remark": "Invalid Field(s): " + strings.Join(validationErrors, ", "),
		})
	}

	logger.LogInfo("Start Tabung process", logrus.Fields{"request": req})

	db := c.Get("db").(*sql.DB)

	nasabah, err := model.GetNasabahByNoRekening(db, req.NoRekening)

	if err != nil {
		logger.LogError("Failed to get nasabah", logrus.Fields{"error": err})
		return c.JSON(http.StatusBadRequest, map[string]string{"remark": err.Error()})
	}

	logger.LogInfo("Cek Saldo process success", logrus.Fields{"request": req})

	return c.JSON(http.StatusOK, map[string]float64{"saldo": nasabah.Saldo})
}
