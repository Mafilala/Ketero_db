
package services

import (
	"context"

	"github.com/Mafilala/ketero/backend/models"
	"github.com/Mafilala/ketero/backend/repositories"
)

func GetClothingByID(ctx context.Context, id int) (*models.Clothing, error) {
	return repositories.FindClothingByID(ctx, id)
}

func CreateNewClothing(ctx context.Context, name string) (*models.Clothing, error) {
	return repositories.CreateNewClothing(ctx, name)
}

func UpdateClothing(ctx context.Context, id int, name string) (*models.Clothing, error) {
	return repositories.UpdateClothing(ctx, id, name)
}

func DeleteClothing(ctx context.Context, id int) (int, error) {
	return repositories.DeleteClothing(ctx, id)
}

func GetAllClothing(ctx context.Context) (*[]models.Clothing, error) {
	return repositories.FindAllClothing(ctx)
}
