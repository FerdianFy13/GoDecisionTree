package model

type AlternativeValueModel struct {
	ID            uint `gorm:"primaryKey"`
	AlternativeID uint
	CriteriaID    uint
	Value         float64
}
