package entity

type Alternative struct {
	ID     int
	Name   string
	Values map[int]float64
	Score  float64
}
