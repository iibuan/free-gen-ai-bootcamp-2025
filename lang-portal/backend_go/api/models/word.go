package models

type Word struct {
	ID              int    `json:"id"`
	BahasaIndonesia string `json:"bahasa_indonesia"`
	English         string `json:"english"`
}
