
package services

import (
    "context"
    "github.com/Mafilala/ketero/backend/repositories"
    "github.com/Mafilala/ketero/backend/models"
    
    )

func GetMeasureByID(ctx context.Context, id int) (*models.Measure, error) {
	return repositories.FindMeasureByID(ctx, id)
}

func CreateNewMeasure(ctx context.Context, name string) (*models.Measure, error) {
	return repositories.CreateNewMeasure(ctx, name)
}

func DeleteMeasure(ctx context.Context, id int) (int, error) {
	return repositories.DeleteMeasure(ctx, id)
}

func GetAllMeasures(ctx context.Context) (*[]models.Measure, error) {
	return repositories.FindAllMeasure(ctx)
}

func UpdateMeasure(ctx context.Context, id int, name string) (*models.Measure, error) {
	return repositories.UpdateMeasure(ctx,id, name)
}
