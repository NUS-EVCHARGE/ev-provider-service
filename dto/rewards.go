package dto

type CoinPolicy struct {
	MaxUsableCoinPerTransaction int  `gorm:"column:max_usable_coin_per_transaction"`
	CashAmount                  int  `gorm:"column:cash_amount"`
	Status                      bool `gorm:"column:status"`
	ProviderId                  int  `gorm:"column:provider_id"`
}

func (CoinPolicy) TableName() string {
	return "coin_tab"
}

type Vouchers struct {
	ProviderId     int    `gorm:"column:provider_id"`
	DiscountAmount int    `gorm:"column:discount_amount"`
	ExpiryDate     string `gorm:"column:expiry_date"`
	Status         bool   `gorm:"column:status"`
	Type           string `gorm:"column:type"`
}

func (Vouchers) TableName() string {
	return "voucher_tab"
}
