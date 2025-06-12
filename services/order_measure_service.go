package services

import (
	"context"
	"github.com/Mafilala/ketero/backend/models"
	"github.com/Mafilala/ketero/backend/repositories"
)

func CreateOrderMeasure(ctx context.Context, om models.OrderMeasure) (*models.OrderMeasure, error) {
	return repositories.CreateOrderMeasure(ctx, om)
}

func UpdateOrderMeasure(ctx context.Context, orderID, clothingID int, measures []models.MeasureValue) error {
	return repositories.UpdateOrderMeasure(ctx, orderID, clothingID, measures)
}

func DeleteOrderMeasure(ctx context.Context, orderID, measureID int) error {
	return repositories.DeleteOrderMeasure(ctx, orderID, measureID)
}

func GetOrderMeasuresByOrderID(ctx context.Context, orderID int) ([]models.OrderMeasure, error) {
	return repositories.GetOrderMeasuresByOrderID(ctx, orderID)
}

