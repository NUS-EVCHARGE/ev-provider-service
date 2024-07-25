package dto

type License struct {
	CompanyId int `gorm:"column:company_id" json:"company_id"`
	Standard  int `gorm:"column:standard" json:"standard"`
	Starter   int `gorm:"column:starter" json:"starter"`
	Premium   int `gorm:"column:premium" json:"premium"`
}

func (License) TableName() string {
	return "license_tab"
}
