package model

type CriteriaModel struct {
	ID     uint   `gorm:"primaryKey"`
	Name   string
	Weight float64
	Type   string
}