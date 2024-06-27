package dto

type ChargerApi struct {
	//gorm.Model
	ID                  uint   `gorm:"primaryKey" json:"id"`
	ProviderId          uint   `gorm:"column:provider_id" json:"provider_id"`
	GetChargerStatusApi string `gorm:"column:get_charger_status_api"`
	StartChargingApi    string `gorm:"column:start_charging_api"`
	EndChargingApi      string `gorm:"column:end_charging_api"`
	HealthCheckApi      string `gorm:"column:health_check_api"`
}
