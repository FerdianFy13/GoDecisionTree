package repository

import "GoDecisionTree/internal/domain/entity"

type AlternativeMemoryRepository struct{}

func NewAlternativeMemoryRepository() *AlternativeMemoryRepository {
	return &AlternativeMemoryRepository{}
}

func (r *AlternativeMemoryRepository) GetDecisionMatrix() (*entity.DecisionMatrix, error) {

	criteria := []entity.Criteria{
		{ID: 1, Name: "Harga", Weight: 0.4, Type: entity.Cost},
		{ID: 2, Name: "Durasi", Weight: 0.3, Type: entity.Cost},
		{ID: 3, Name: "Bagasi", Weight: 0.2, Type: entity.Benefit},
		{ID: 4, Name: "Rating", Weight: 0.1, Type: entity.Benefit},
	}

	alternatives := []entity.Alternative{
		{
			ID:   1,
			Name: "Garuda",
			Values: map[int]float64{
				1: 1500000,
				2: 2,
				3: 20,
				4: 4.8,
			},
		},
		{
			ID:   2,
			Name: "Lion",
			Values: map[int]float64{
				1: 900000,
				2: 2.5,
				3: 15,
				4: 3.8,
			},
		},
		{
			ID:   3,
			Name: "AirAsia",
			Values: map[int]float64{
				1: 850000,
				2: 3,
				3: 10,
				4: 4.0,
			},
		},
	}

	return &entity.DecisionMatrix{
		Criteria:     criteria,
		Alternatives: alternatives,
	}, nil
}
