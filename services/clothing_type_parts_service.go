
package services

import (
	"context"

	"github.com/Mafilala/ketero/backend/models"
	"github.com/Mafilala/ketero/backend/repositories"
)

func AddClothing(ctx context.Context, clothing_type_id, clothing_id int) (*models.ClothingTypePart, error) {
	return repositories.AddClothing(ctx, clothing_type_id, clothing_id)
}

func RemoveClothingTypePart(ctx context.Context, clothing_type_id, clothing_id int) (int, error) {
	return repositories.RemoveClothingTypePart(ctx, clothing_type_id, clothing_id)
}

func GetAllClothingParts(ctx context.Context, clothing_type_id int) ([]models.Clothing, error) {
	return repositories.GetAllClothingPart(ctx, clothing_type_id)
} 
