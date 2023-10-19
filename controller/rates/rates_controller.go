package rates

import (
	"github.com/NUS-EVCHARGE/ev-provider-service/dao"
	"github.com/NUS-EVCHARGE/ev-provider-service/dto"
)

type RateController interface {
	AddRate(Rate *dto.Rates) error
	GetRateByProviderId(providerId uint) ([]dto.Rates, error)
	GetRateByRateId(rateId uint) (dto.Rates, error)
	DeleteRate(RateId uint) error
	UpdateRate(Rate dto.Rates) error
}

type RateControllerImpl struct {
}

func (r *RateControllerImpl) AddRate(Rate *dto.Rates) error {
	err := dao.Db.CreateRatesEntry(Rate)
	return err
}

func (r *RateControllerImpl) GetRateByProviderId(providerId uint) ([]dto.Rates, error) {
	return dao.Db.GetRatesByProviderId(providerId)
}

func (r *RateControllerImpl) GetRateByRateId(rateId uint) (dto.Rates, error) {
	return dao.Db.GetRatesByRateId(rateId)
}

func (r *RateControllerImpl) DeleteRate(RateId uint) error {
	return dao.Db.DeleteRatesEntry(dto.Rates{ID: RateId})
}

func (r *RateControllerImpl) UpdateRate(Rate dto.Rates) error {
	return dao.Db.UpdateRatesEntry(Rate)
}

var (
	RateControllerObj RateController
)

func NewRateController() {
	RateControllerObj = &RateControllerImpl{}
}
