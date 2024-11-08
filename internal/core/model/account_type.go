package model

import "time"

type AccountType struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Enum      string    `gorm:"column:enum;type:varchar(255);unique" json:"enum"`
	NameTH    string    `gorm:"column:name_th;type:varchar(255)" json:"name_th"`
	NameEN    string    `gorm:"column:name_en;type:varchar(255)" json:"name_en"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
