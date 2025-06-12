
package services

import (
    "context"
    "github.com/Mafilala/ketero/backend/models"
    "github.com/Mafilala/ketero/backend/repositories"
)

func GetStatusByID(ctx context.Context, id int) (*models.Status, error) {
    return repositories.FindStatusByID(ctx, id)
}

func CreateNewStatus(ctx context.Context, name string) (*models.Status, error) {
    return repositories.CreateNewStatus(ctx, name)
}

func DeleteStatus(ctx context.Context, id int) (int, error) {
    return repositories.DeleteStatus(ctx, id)
}

func GetAllStatuses(ctx context.Context) (*[]models.Status, error) {
    return repositories.FindAllStatuses(ctx)
}

func UpdateStatus(ctx context.Context, id int, name string) (*models.Status, error) {
	return repositories.UpdateStatus(ctx, id, name)
}
