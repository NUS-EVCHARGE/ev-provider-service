package dto

type Rates struct {
	ID            uint    `gorm:"primaryKey" json:"id"`
	ProviderId    uint    `gorm:"column:provider_id"`
	NormalRate    float64 `gorm:"column:normal_rate"`
	PenaltyRate   float64 `gorm:"column:penalty_rate"`
	NoShowPenalty float64 `gorm:"column:no_show_penalty"`
	Status        string  `gorm:"column:status"`
}

func (Rates) TableName() string {
	return "rates_tab"
}
