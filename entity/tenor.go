package entity

type Tenor struct {
	ID          uint    `json:"tenor_id" gorm:"primaryKey"`
	CustomerID  string  `json:"customer_id" gorm:"not null"`
	Limit       float64 `json:"limit"`
	MonthNumber int     `json:"month_number"`
	IsLunas     bool    `json:"is_lunas"`
}
