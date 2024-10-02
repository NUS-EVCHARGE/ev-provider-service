package dto

type HeatMapData struct {
	Lat         float64 `csv:"latitude" json:"lat"`
	Lng         float64 `csv:"longitude" json:"lng"`
	DemandScore int     `csv:"demand_score" json:"demand_score"`
}
