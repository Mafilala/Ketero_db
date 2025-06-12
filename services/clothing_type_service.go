
package services

import (
    "context"
    "github.com/Mafilala/ketero/backend/repositories"
    "github.com/Mafilala/ketero/backend/models"
    
    )

func GetClothingTypeByID(ctx context.Context, id int) (*models.ClothingType, error) {
	return repositories.FindClothingTypeByID(ctx, id)
}

func CreateNewClothingType(ctx context.Context, name string) (*models.ClothingType, error) {
	return repositories.CreateNewClothingType(ctx, name)
}

func DeleteClothingType(ctx context.Context, id int) (int, error) {
	return repositories.DeleteClothingType(ctx, id)
}

func GetAllClothingTypes(ctx context.Context) (*[]models.ClothingType, error) {
	return repositories.FindAllClothingType(ctx)
}

func UpdateClothingType(ctx context.Context, id int, name string) (*models.ClothingType, error) {
	return repositories.UpdateClothingType(ctx, id, name)
} 
