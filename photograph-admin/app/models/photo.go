package models

import "time"

type Photo struct {
	Base
	FileName     string    `gorm:"size:225" json:"fileName,omitempty"`
	Name         string    `gorm:"size:225" json:"name,omitempty"`
	Location     string    `gorm:"size:225" json:"location,omitempty"`
	Color        uint      `gorm:"size:500" json:"color,omitempty"`
	Genre        uint      `json:"genre,omitempty"`
	Shootingdate time.Time `json:"shooting_date"`
}
