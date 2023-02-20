package models

import "time"

type Paket struct {
	PaketID     int        `json:"paket_id"`
	OwnerID     int        `json:"owner_id"`
	PaketName   string     `json:"paket_name"`
	Desc        string     `json:"desc"`
	DurasiStart int        `json:"durasi_start"`
	DurasiEnd   int        `json:"durasi_end"`
	Price       float32    `json:"price"`
	IsActive    bool       `json:"is_active"`
	IsPromo     bool       `json:"is_promo"`
	PromoPrice  float32    `json:"promo_price"`
	PromoStart  *time.Time `json:"promo_start"`
	PromoEnd    *time.Time `json:"promo_end"`
}
