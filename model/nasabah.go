package model

import (
	"database/sql"
	"fmt"
	"math/rand/v2"
	"strconv"

	logger "service-account/logger"

	"github.com/sirupsen/logrus"

	_ "github.com/lib/pq"
)

type Nasabah struct {
	ID         int     `json:"id"`
	Nama       string  `json:"nama"`
	NIK        string  `json:"nik"`
	NoHP       string  `json:"no_hp"`
	NoRekening string  `json:"no_rekening"`
	Saldo      float64 `json:"saldo"`
}

func CreateNasabah(db *sql.DB, nama, nik, noHP string) (string, error) {
	noRekening := createNoRekening()
	for IsDuplicateNoRek(db, noRekening) {
		noRekening = createNoRekening()
	}

	// Insert ke database
	_, err := db.Exec(`INSERT INTO nasabah (nama, nik, no_hp, no_rekening) VALUES ($1, $2, $3, $4)`,
		nama, nik, noHP, noRekening)
	if err != nil {
		logger.LogError("Error while inserting nasabah: %v", logrus.Fields{"error": err, "data": map[string]interface{}{
			"nama":       nama,
			"nik":        nik,
			"noHP":       noHP,
			"noRekening": noRekening,
		}})
		return "", err
	}

	logger.LogInfo("Nasabah successfully created.", logrus.Fields{"nama": nama, "noRekening": noRekening})
	return strconv.Itoa(noRekening), nil
}

func createNoRekening() int {
	return rand.IntN(999999999999999)
}

func IsDuplicateNoRek(db *sql.DB, noRekening int) bool {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM nasabah WHERE no_rekening = $1", noRekening).Scan(&count)
	if err != nil {
		logger.LogError("Error checking No Rekening: %v", logrus.Fields{"error": err})
		return false
	}
	return count > 0
}

func IsDuplicateNIK(db *sql.DB, nik string) bool {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM nasabah WHERE nik = $1", nik).Scan(&count)
	if err != nil {
		logger.LogError("Error checking NIK: %v", logrus.Fields{"error": err})
		return false
	}
	return count > 0
}

func IsDuplicateNoHP(db *sql.DB, noHP string) bool {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM nasabah WHERE no_hp = $1", noHP).Scan(&count)
	if err != nil {
		logger.LogError("Error checking No HP: %v", logrus.Fields{"error": err})
		return false
	}
	return count > 0
}

func GetNasabahByNoRekening(db *sql.DB, noRekening string) (Nasabah, error) {
	var nasabah Nasabah
	err := db.QueryRow("SELECT id, nama, nik, no_hp, no_rekening, saldo FROM nasabah WHERE no_rekening = $1", noRekening).
		Scan(&nasabah.ID, &nasabah.Nama, &nasabah.NIK, &nasabah.NoHP, &nasabah.NoRekening, &nasabah.Saldo)
	if err != nil {
		if err == sql.ErrNoRows {
			return Nasabah{}, fmt.Errorf("no rekening tidak ditemukan")
		}
		logger.LogError("Error fetching nasabah by no rekening", logrus.Fields{"error": err, "data": map[string]interface{}{
			"noRekening": noRekening,
		}})
		return Nasabah{}, err
	}
	return nasabah, nil
}

func UpdateSaldoByNoRekening(db *sql.DB, noRekening string, newSaldo, oldSaldo float64) (float64, error) {
	_, err := db.Exec("UPDATE nasabah SET saldo = $2 WHERE no_rekening = $1;", noRekening, newSaldo)
	if err != nil {
		logger.LogError("Error Update Saldo", logrus.Fields{"error": err, "data": map[string]interface{}{
			"noRekening": noRekening,
			"newSaldo":   newSaldo,
			"oldSaldo":   oldSaldo,
		}})
		return 0, err
	}
	return newSaldo, nil
}
