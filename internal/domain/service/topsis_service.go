package service

import (
	"GoDecisionTree/internal/domain/entity"
	"GoDecisionTree/internal/domain/repository"
	"math"
	"sort"
)

type TopsisService struct {
	repo repository.AlternativeRepository
}

func NewTopsisService(repo repository.AlternativeRepository) *TopsisService {
	return &TopsisService{repo: repo}
}

func (s *TopsisService) Rank() ([]entity.Alternative, error) {
	matrix, err := s.repo.GetDecisionMatrix()
	if err != nil {
		return nil, err
	}

	criteria := matrix.Criteria
	alternatives := matrix.Alternatives

	// Step 1: Normalization
	denominator := make(map[int]float64)

	for _, c := range criteria {
		for _, a := range alternatives {
			denominator[c.ID] += math.Pow(a.Values[c.ID], 2)
		}
		denominator[c.ID] = math.Sqrt(denominator[c.ID])
	}

	normalized := make([]entity.Alternative, len(alternatives))

	for i, a := range alternatives {
		values := make(map[int]float64)
		for _, c := range criteria {
			values[c.ID] = (a.Values[c.ID] / denominator[c.ID]) * c.Weight
		}
		normalized[i] = entity.Alternative{
			ID:     a.ID,
			Name:   a.Name,
			Values: values,
		}
	}

	// Step 2: Ideal Solution
	idealPositive := make(map[int]float64)
	idealNegative := make(map[int]float64)

	for _, c := range criteria {
		first := normalized[0].Values[c.ID]
		idealPositive[c.ID] = first
		idealNegative[c.ID] = first

		for _, a := range normalized {
			val := a.Values[c.ID]
			if c.Type == entity.Benefit {
				idealPositive[c.ID] = math.Max(idealPositive[c.ID], val)
				idealNegative[c.ID] = math.Min(idealNegative[c.ID], val)
			} else {
				idealPositive[c.ID] = math.Min(idealPositive[c.ID], val)
				idealNegative[c.ID] = math.Max(idealNegative[c.ID], val)
			}
		}
	}

	// Step 3: Distance & Score
	for i := range normalized {
		var dPlus, dMinus float64

		for _, c := range criteria {
			val := normalized[i].Values[c.ID]
			dPlus += math.Pow(val-idealPositive[c.ID], 2)
			dMinus += math.Pow(val-idealNegative[c.ID], 2)
		}

		dPlus = math.Sqrt(dPlus)
		dMinus = math.Sqrt(dMinus)

		normalized[i].Score = dMinus / (dPlus + dMinus)
	}

	sort.Slice(normalized, func(i, j int) bool {
		return normalized[i].Score > normalized[j].Score
	})

	return normalized, nil
}
