package dto

type ChargerPoint struct {
	//gorm.Model
	ID         uint    `gorm:"primaryKey" json:"id"`
	ProviderId uint    `gorm:"column:provider_id" json:"provider_id"`
	Lat        float64 `gorm:"column:lat" json:"lat"`
	Lng        float64 `gorm:"column:lng" json:"lng"`
	Address    string  `gorm:"colummn:address" json:"address"`
}

type Charger struct {
	ID             uint    `gorm:"primaryKey" json:"id"`
	ChargerPointID uint    `gorm:"column:charger_point_id" json:"charger_point_id"`
	UID            string  `gorm:"column:uid" json:"uid"`
	Status         string  `gorm:"column:status" json:"status"`
	PowerType      string  `gorm:"column:power_type" json:"power_type"`
	ChargerType    string  `gorm:"column:charger_type" json:"charger_type"`
	Power          float64 `gorm:"colum:power" json:"power"`
	Details        string  `gorm:"column:details" json:"details"` // json struct
}

type ChargerDetails struct {
	ChargerType string  `json:"charger_type"`
	Rates       float64 `json:"rates"`
	// todo: include href to start charging
}

func (Charger) TableName() string {
	return "charger_tab"
}

func (ChargerPoint) TableName() string {
	return "charger_point_tab"
}

type IoTHubRequest struct {
	Command   string `json:"command"`
	CompanyId string `json:"company_id"`
	ChargerId string `json:"charger_id"`
	Status    string `json:"status"`
}
type ChargerFullDetails struct {
	Key            string  `json:"key"`
	ProviderName   string  `json:"provider_name"`
	Lat            float64 `json:"lat"`
	Lng            float64 `json:"lng"`
	Address        string  `json:"address"`
	ID             uint    `gorm:"primaryKey" json:"id"`
	ChargerPointID uint    `gorm:"column:charger_point_id" json:"charger_point_id"`
	UID            string  `gorm:"column:uid" json:"uid"`
	Status         string  `gorm:"column:status" json:"status"`
	PowerType      string  `gorm:"column:power_type" json:"power_type"`
	ChargerType    string  `gorm:"column:charger_type" json:"charger_type"`
	Power          float64 `gorm:"colum:power" json:"power"`
	Details        string  `gorm:"column:details" json:"details"` // json struct
	// to be removed
	ChargerList []Charger `json:"charger_list"`
}
