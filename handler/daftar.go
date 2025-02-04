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

func Daftar(c echo.Context) error {
	var req struct {
		Nama string `json:"nama" validate:"required,min=3,max=100"`
		NIK  string `json:"nik" validate:"required,len=16"`
		NoHP string `json:"no_hp" validate:"required,len=12,number"`
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
	logger.LogInfo("Start Daftar process", logrus.Fields{"request": req})

	db := c.Get("db").(*sql.DB)

	if model.IsDuplicateNIK(db, req.NIK) {
		logger.LogError("Duplicate NIK", logrus.Fields{"NIK": req.NIK})
		return c.JSON(http.StatusBadRequest, map[string]string{"remark": "NIK sudah digunakan"})
	}
	if model.IsDuplicateNoHP(db, req.NoHP) {
		logger.LogError("Duplicate NoHP", logrus.Fields{"NoHP": req.NoHP})
		return c.JSON(http.StatusBadRequest, map[string]string{"remark": "No HP sudah digunakan"})
	}

	noRekening, err := model.CreateNasabah(db, req.Nama, req.NIK, req.NoHP)
	if err != nil {
		logger.LogError("Failed to create nasabah", logrus.Fields{"error": err})
		return c.JSON(http.StatusInternalServerError, map[string]string{"remark": "Gagal membuat nasabah"})
	}

	logger.LogInfo("Daftar process success", logrus.Fields{"request": req})

	return c.JSON(http.StatusOK, map[string]string{"no_rekening": noRekening})
}
