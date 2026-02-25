package test

import (
	"GoDecisionTree/internal/domain/entity"
	"GoDecisionTree/internal/domain/service"
	"testing"
)

type MockRepo struct{}

func (m *MockRepo) GetDecisionMatrix() (*entity.DecisionMatrix, error) {

	return &entity.DecisionMatrix{
		Criteria: []entity.Criteria{
			{ID: 1, Weight: 1, Type: entity.Benefit},
		},
		Alternatives: []entity.Alternative{
			{ID: 1, Name: "A", Values: map[int]float64{1: 10}},
			{ID: 2, Name: "B", Values: map[int]float64{1: 5}},
		},
	}, nil
}

func TestTopsisRanking(t *testing.T) {

	service := service.NewTopsisService(&MockRepo{})
	result, err := service.Rank()

	if err != nil {
		t.Fatal(err)
	}

	if result[0].Name != "A" {
		t.Errorf("Expected A as best alternative")
	}
}