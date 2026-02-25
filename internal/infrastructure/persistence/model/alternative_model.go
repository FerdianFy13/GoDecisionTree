package model

type AlternativeModel struct {
	ID     uint `gorm:"primaryKey"`
	Name   string
	Values []AlternativeValueModel `gorm:"foreignKey:AlternativeID"`
}
