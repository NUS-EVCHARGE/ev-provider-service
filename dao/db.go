package dao

import (
	"github.com/NUS-EVCHARGE/ev-provider-service/dto"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database interface {
	// provider impl
	CreateProviderEntry(Provider dto.Provider) (dto.Provider, error)
	UpdateProviderEntry(Provider dto.Provider) error
	DeleteProviderEntry(Provider dto.Provider) error
	GetProviderEntry(email string) (dto.Provider, error)
	GetProviderEntryByCompany(companyName string) (dto.Provider, error)
	GetAllProviderEntry() ([]dto.Provider, error)
	// charger point impl
	CreateChargerPointEntry(chargerPoint *dto.ChargerPoint) error
	GetChargerPointEntryByID(chargerId uint) (dto.ChargerPoint, error)
	GetChargerPointEntryByProviderID(providerId uint) ([]dto.ChargerPoint, error)
	GetAllChargerPointEntry() ([]dto.ChargerPoint, error)
	GetAllChargerPointEntryByProviderID(providerId int) ([]dto.ChargerPoint, error)
	UpdateChargerPointEntry(chargerPoint dto.ChargerPoint) error
	// charger impl
	CreateChargerEntry(charger dto.Charger) error
	GetChargerPointByLocation(providerId int, placeId string) (dto.ChargerPoint, error)
	GetChargerByChargerPointId(chargerPointId uint) ([]dto.Charger, error)
	GetChargerById(chargerId uint) (dto.Charger, error)
	UpdateChargerEntry(charger dto.Charger) error
	// license impl
	CreateLicense(license dto.License) (dto.License, error)
	GetLicenseByCompanyId(companyId int) (dto.License, error)
	UpdateLicense(license dto.License) error
	// rewards impl
	CreateCoinPolicy(coinPolicy dto.CoinPolicy) (dto.CoinPolicy, error)
	UpdateCoinPolicy(coinPolicy dto.CoinPolicy) (dto.CoinPolicy, error)
	GetCoinPolicy(providerId int) (dto.CoinPolicy, error)
	CreateVouchers(voucher dto.Vouchers) (dto.Vouchers, error)
	GetAllVouchers(providerId int) ([]dto.Vouchers, error)
	UpdateVoucher(voucher dto.Vouchers) (dto.Vouchers, error)
	SetVouchersToBeExpired(currentTime int64) error
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
