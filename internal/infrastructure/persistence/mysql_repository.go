package persistence

import (
	"GoDecisionTree/internal/domain/entity"
	"GoDecisionTree/internal/domain/repository"
	"GoDecisionTree/internal/infrastructure/persistence/model"

	"gorm.io/gorm"
)

type MySQLRepository struct {
	db *gorm.DB
}

func NewMySQLRepository(db *gorm.DB) repository.AlternativeRepository {
	return &MySQLRepository{db: db}
}

func (r *MySQLRepository) GetDecisionMatrix() (*entity.DecisionMatrix, error) {

	var criteriaModels []model.CriteriaModel
	var alternativeModels []model.AlternativeModel

	r.db.Find(&criteriaModels)
	r.db.Preload("Values").Find(&alternativeModels)

	var criteria []entity.Criteria
	for _, c := range criteriaModels {
		criteria = append(criteria, entity.Criteria{
			ID:     int(c.ID),
			Name:   c.Name,
			Weight: c.Weight,
			Type:   entity.CriteriaType(c.Type),
		})
	}

	var alternatives []entity.Alternative
	for _, a := range alternativeModels {

		values := make(map[int]float64)
		for _, v := range a.Values {
			values[int(v.CriteriaID)] = v.Value
		}

		alternatives = append(alternatives, entity.Alternative{
			ID:     int(a.ID),
			Name:   a.Name,
			Values: values,
		})
	}

	return &entity.DecisionMatrix{
		Criteria:     criteria,
		Alternatives: alternatives,
	}, nil
}
