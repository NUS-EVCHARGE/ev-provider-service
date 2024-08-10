package dto

type CoinPolicy struct {
	ID                          int   `gorm:"column:id"`
	MaxUsableCoinPerTransaction int   `gorm:"column:max_usable_coin_per_transaction" json:"max_usable_coin_per_transaction"`
	CashAmount                  int   `gorm:"column:cash_amount" json:"cash_amount"`
	Status                      *bool `gorm:"column:status" json:"status"`
	ProviderId                  int   `gorm:"column:provider_id" json:"provider_id"`
}

func (CoinPolicy) TableName() string {
	return "coin_tab"
}

type Vouchers struct {
	ID               int    `gorm:"column:id" json:"id"`
	Name             string `gorm:"column:name" json:"name"`
	ProviderId       int    `gorm:"column:provider_id"`
	DiscountAmount   int    `gorm:"column:discount_amount" json:"discount_amount"`
	ExpiryDate       string `json:"expiry_date" gorm:"-"`
	ExpiryDateInUnix int64  `gorm:"column:expiry_date"` // in unix
	Status           bool   `gorm:"column:status"`
	Type             string `gorm:"column:type" json:"type"`
}

func (Vouchers) TableName() string {
	return "voucher_tab"
}
