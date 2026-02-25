package repository

import "GoDecisionTree/internal/domain/entity"

type AlternativeRepository interface {
	GetDecisionMatrix() (*entity.DecisionMatrix, error)
}