package mapper

import (
	"GoDecisionTree/internal/application/dto"
	"GoDecisionTree/internal/domain/entity"
)

func ToRankingResponse(alternatives []entity.Alternative) []dto.RankingResponse {

	var result []dto.RankingResponse

	for _, a := range alternatives {
		result = append(result, dto.RankingResponse{
			ID:    a.ID,
			Name:  a.Name,
			Score: a.Score,
		})
	}

	return result
}