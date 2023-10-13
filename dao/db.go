package dao

import (
	"github.com/NUS-EVCHARGE/ev-provider-service/dto"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database interface {
	CreateProviderEntry(Provider dto.Provider) (dto.Provider, error)
	UpdateProviderEntry(Provider dto.Provider) error
	DeleteProviderEntry(Provider dto.Provider) error
	GetAllProviderEntry(email string) (dto.Provider, error)

	CreateChargerEntry(Charger dto.Charger) error
	UpdateChargerEntry(Charger dto.Charger) error
	DeleteChargerEntry(Charger dto.Charger) error
	GetAllChargerEntry(providerId uint) ([]dto.Charger, error)
	GetChargerById(chargerId uint) (dto.Charger, error)

	CreateRatesEntry(Rates dto.Rates) error
	UpdateRatesEntry(Rates dto.Rates) error
	DeleteRatesEntry(Rates dto.Rates) error
	GetRatesByProviderId(providerId uint) ([]dto.Rates, error)
	GetRatesByRateId(rateId uint) (dto.Rates, error)
}

var (
	Db Database
)

type dbImpl struct {
	Dsn          string
	DbController *gorm.DB
}

func InitDB(dsn string) error {
	if dbObj, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		return err
	} else {
		Db = NewDatabase(dsn, dbObj)
		return nil
	}
}

func NewDatabase(dsn string, dbObj *gorm.DB) Database {
	return &dbImpl{
		Dsn:          dsn,
		DbController: dbObj,
	}
}
