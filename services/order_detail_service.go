
package services

import (
	"context"
	"github.com/Mafilala/ketero/backend/models"
	"github.com/Mafilala/ketero/backend/repositories"
)

func CreateOrderDetail(ctx context.Context, detail models.OrderDetail) (*models.OrderDetail, error) {
	return repositories.CreateOrderDetail(ctx, detail)
}

func UpdateOrderDetail(ctx context.Context, detail models.OrderDetail) error {
	return repositories.UpdateOrderDetail(ctx, detail)
}

func DeleteOrderDetail(ctx context.Context, orderID int) error {
	return repositories.DeleteOrderDetail(ctx, orderID)
}

func GetOrderDetail(ctx context.Context, orderID int) (*models.OrderDetail, error) {
	return repositories.GetOrderDetail(ctx, orderID)
}
