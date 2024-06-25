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
	GetProviderEntry(email string) (dto.Provider, error)
	GetProviderEntryByCompany(companyName string) (dto.Provider, error)
	GetAllProviderEntry() ([]dto.Provider, error)

	CreateChargerPointEntry(chargerPoint *dto.ChargerPoint) error
	GetChargerPointEntryByID(chargerId uint) (dto.ChargerPoint, error)
	GetChargerPointEntryByProviderID(providerId uint) ([]dto.ChargerPoint, error)
	GetAllChargerPointEntry() ([]dto.ChargerPoint, error)
	GetAllChargerPointEntryByProviderID(providerId int) ([]dto.ChargerPoint, error)
	UpdateChargerPointEntry(chargerPoint dto.ChargerPoint) error

	CreateChargerEntry(charger dto.Charger) error
	GetChargerPointByLocation(providerId int, lat, lng float64) (dto.ChargerPoint, error)
	GetChargerByChargerPointId(chargerPointId uint) ([]dto.Charger, error)
	GetChargerById(chargerId uint) (dto.Charger, error)
	UpdateChargerEntry(charger dto.Charger) error
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
