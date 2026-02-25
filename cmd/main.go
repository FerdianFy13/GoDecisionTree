package main

import (
	"fmt"
	"GoDecisionTree/internal/application/mapper"
	"GoDecisionTree/internal/domain/service"
	"GoDecisionTree/internal/infrastructure/config"
	"GoDecisionTree/internal/infrastructure/persistence"
)

func main() {

	db, _ := config.NewDatabase()

	repo := persistence.NewMySQLRepository(db)
	topsisService := service.NewTopsisService(repo)

	result, _ := topsisService.Rank()

	response := mapper.ToRankingResponse(result)

	for i, r := range response {
		fmt.Printf("%d. %s (%.4f)\n", i+1, r.Name, r.Score)
	}
}