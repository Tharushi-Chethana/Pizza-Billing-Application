package models

type InvoiceItem struct {
	InvoiceItemID        uint    `gorm:"primaryKey" json:"id"`
	InvoiceID uint    `json:"invoice_id"`
	ItemID    uint    `json:"item_id"`
	Quantity  int     `json:"quantity"`
	Total     float64 `json:"total"`
}