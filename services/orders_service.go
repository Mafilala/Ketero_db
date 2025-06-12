
package services

import (
    "context"
    "github.com/Mafilala/ketero/backend/models"
    "github.com/Mafilala/ketero/backend/repositories"
)

func CreateNewOrder(ctx context.Context, order models.Order) (*models.Order, error) {
    return repositories.CreateOrder(ctx, order)
}

func GetOrderByID(ctx context.Context, id int) (*models.Order, error) {
    return repositories.FindOrderByID(ctx, id)
}

func DeleteOrder(ctx context.Context, id int) (int, error) {
    return repositories.DeleteOrder(ctx, id)
}

func GetAllOrders(ctx context.Context, limit, offset int, status string) ([]models.Order, int, error) {
    return repositories.FindAllOrders(ctx, limit, offset, status)
}

func PatchOrder(ctx context.Context, id int, patch *models.Order)(*models.Order, error) {
	return repositories.PatchOrder(ctx, id, patch)
}
