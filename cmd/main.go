package main

import (
	"fmt"
	"GoDecisionTree/internal/domain/service"
	infra "GoDecisionTree/internal/infrastructure/repository"
)

func main() {
	repo := infra.NewAlternativeMemoryRepository()
	service := service.NewTopsisService(repo)

	result, _ := service.Rank()

	fmt.Println("=== Ranking Tiket Terbaik ===")
	for i, r := range result {
		fmt.Printf("%d. %s (Score: %.4f)\n", i+1, r.Name, r.Score)
	}
}
