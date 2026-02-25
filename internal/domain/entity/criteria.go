package entity

type CriteriaType string

const (
	Benefit CriteriaType = "benefit"
	Cost    CriteriaType = "cost"
)

type Criteria struct {
	ID     int
	Name   string
	Weight float64
	Type   CriteriaType
}