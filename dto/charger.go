package dto

type Charger struct {
	//gorm.Model
	ID         uint    `gorm:"primaryKey" json:"id"`
	ProviderId uint    `gorm:"column:provider_id"`
	RatesId    uint    `gorm:"column:rates_id"`
	Lat        float64 `gorm:"column:lat"`
	Lng        float64 `gorm:"column:lng"`
	Status     string  `gorm:"column:status"`
}

func (Charger) TableName() string {
	return "charger_tab"
}
