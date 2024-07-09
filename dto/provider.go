package dto

type Provider struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	UserEmail   string `gorm:"column:user_email" json:"email"`
	CompanyName string `gorm:"column:company_name" json:"company_name"`
	Description string `gorm:"column:description" json:"description"`
	Status      string `gorm:"column:status" json:"status"`
}

func (Provider) TableName() string {
	return "provider_tab"
}
