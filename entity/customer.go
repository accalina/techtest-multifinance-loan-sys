package entity

import "time"

type DetailCustomer struct {
	NIK          string    `json:"nik" gorm:"primaryKey"`
	FullName     string    `json:"full_name"`
	LegalName    string    `json:"legal_name"`
	TempatLahir  string    `json:"tempat_lahir"`
	TanggalLahir time.Time `json:"tanggal_lahir"`
	Gaji         float64   `json:"gaji"`
	FotoKTP      string    `json:"foto_ktp"`
	FotoSelfie   string    `json:"foto_selfie"`
}

type CustomerPayload struct {
	NIK          string  `json:"nik"`
	FullName     string  `json:"full_name"`
	LegalName    string  `json:"legal_name"`
	TempatLahir  string  `json:"tempat_lahir"`
	TanggalLahir string  `json:"tanggal_lahir"` // Date as string from JSON
	Gaji         float64 `json:"gaji"`
	FotoKTP      string  `json:"foto_ktp"`
	FotoSelfie   string  `json:"foto_selfie"`
}
