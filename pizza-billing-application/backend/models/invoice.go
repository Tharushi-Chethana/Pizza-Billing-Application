package models

import "time"

type Invoice struct {
	InvoiceID          uint          `gorm:"primaryKey" json:"id"`
	Customer    string        `json:"customer"`
	CreatedAt   time.Time     `json:"created_at"`
	TotalAmount float64       `json:"total_amount"`
	Tax         float64       `json:"tax"`
	GrandTotal  float64       `json:"grand_total"`
	Items       []InvoiceItem `gorm:"foreignKey:InvoiceID" json:"items"`
}