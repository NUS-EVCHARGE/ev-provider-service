package dto

type Charger struct {
	//gorm.Model
	ID         uint    `gorm:"primaryKey" json:"id"`
	ProviderId uint    `gorm:"column:provider_id" json:"provider_id"`
	RatesId    uint    `gorm:"column:rates_id" json:"rates_id"`
	Address    string  `gorm:"address" json:"address"`
	Lat        float64 `gorm:"column:lat" json:"lat"`
	Lng        float64 `gorm:"column:lng" json:"lng"`
	Status     string  `gorm:"column:status" json:"status"`
}

func (Charger) TableName() string {
	return "charger_tab"
}
