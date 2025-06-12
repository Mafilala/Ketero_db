
package services

import (
	"context"

	"github.com/Mafilala/ketero/backend/models"
	"github.com/Mafilala/ketero/backend/repositories"
)

func AddMeasure(ctx context.Context, clothing_id, measure_id int) (*models.ClothingMeasures, error) {
	return repositories.AddMeasure(ctx, clothing_id, measure_id)
}

func RemoveMeasure(ctx context.Context, clothing_id, measure_id int) (int, error) {
	return repositories.RemoveMeasure(ctx, clothing_id, measure_id)
}

func GetAllClothingMeasures(ctx context.Context, clothing_id int) (*[]models.Measure, error) {
	return repositories.GetAllClothingMeasures(ctx, clothing_id)
} 
