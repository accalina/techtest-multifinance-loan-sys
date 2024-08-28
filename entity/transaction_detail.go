package entity

type TransactionDetail struct {
	ID                uint    `json:"transaction_id" gorm:"primaryKey"`
	CustomerID        string  `json:"customer_id" gorm:"not null"`
	OTRPrice          float64 `json:"otr_price"`
	AdminFee          float64 `json:"admin_fee"`
	InstallmentAmount float64 `json:"installment_amount"`
	InterestAmount    float64 `json:"interest_amount"`
	AssetName         string  `json:"asset_name"`
}
