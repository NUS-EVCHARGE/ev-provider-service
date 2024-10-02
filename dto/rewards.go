package dto

type CoinPolicy struct {
	ID                          int    `gorm:"column:id"`
	Name                        string `gorm:"column:name" json:"policy_name"`
	MaxUsableCoinPerTransaction int    `gorm:"column:max_usable_coin_per_transaction" json:"max_usable_coin_per_transaction"`
	CashAmount                  int    `gorm:"column:cash_amount" json:"cash_amount"`
	Status                      *bool  `gorm:"column:status" json:"status"`
	ProviderId                  int    `gorm:"column:provider_id" json:"provider_id"`
}

func (CoinPolicy) TableName() string {
	return "coin_tab"
}

type Vouchers struct {
	ID              int    `gorm:"column:id" json:"id"`
	Key             int    `gorm:"-" json:"key"`
	Name            string `gorm:"column:name" json:"name"`
	ProviderId      int    `gorm:"column:provider_id"`
	DiscountAmount  int    `gorm:"column:discount_amount" json:"discount_amount"`
	StartDate       string `json:"start_date" gorm:"-"`
	StartDateInUnix int64  `gorm:"column:start_date"` // in unix
	EndDate         string `json:"end_date" gorm:"-"`
	EndDateInUnix   int64  `gorm:"column:end_date"` // in unix
	Status          bool   `gorm:"column:status" json:"status"`
	Quota           int    `gorm:"column:quota" json:"quota"`
	RedeemedAmount  int    `gorm:"redeemed_amount" json:"redeemed_amount"`
	QuotaUnit       string `gorm:"quota_unit" json:"quota_unit"` // user , global
	VoucherType     string `gorm:"column:voucher_type" json:"voucher_type"`
	DiscountType    string `gorm:"column:discount_type" json:"discount_type"`
}

func (Vouchers) TableName() string {
	return "voucher_tab"
}
