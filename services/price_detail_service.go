
package services

import (
	"context"
	"github.com/Mafilala/ketero/backend/models"
	"github.com/Mafilala/ketero/backend/repositories"
)

func CreatePriceDetail(ctx context.Context, detail models.PriceDetail) (*models.PriceDetail, error) {
	return repositories.CreatePriceDetail(ctx, detail)
}

func GetPriceDetailByOrderID(ctx context.Context, orderID int) (*models.PriceDetail, error) {
	return repositories.GetPriceDetailByOrderID(ctx, orderID)
}

func UpdatePriceDetail(ctx context.Context, detail models.PriceDetail) (*models.PriceDetail, error) {
	return repositories.UpdatePriceDetail(ctx, detail)
}

func DeletePriceDetail(ctx context.Context, orderID int) error {
	return repositories.DeletePriceDetail(ctx, orderID)
}

