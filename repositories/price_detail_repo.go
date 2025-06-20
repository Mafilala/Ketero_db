
package repositories

import (
	"context"
	"github.com/Mafilala/ketero/backend/initializers"
	"github.com/Mafilala/ketero/backend/models"
)

func CreatePriceDetail(ctx context.Context, detail models.PriceDetail) (*models.PriceDetail, error) {
	query := `INSERT INTO price_detail (order_id, price, paid) VALUES ($1, $2, $3)`
	_, err := initializers.Pool.Exec(ctx, query, detail.OrderID, *detail.Price, *detail.Paid)
	if err != nil {
		return nil, err
	}
	return &detail, nil
}

func GetPriceDetailByOrderID(ctx context.Context, orderID int) (*models.PriceDetail, error) {
	query := `SELECT order_id, price, paid FROM price_detail WHERE order_id = $1`
	var detail models.PriceDetail
	err := initializers.Pool.QueryRow(ctx, query, orderID).Scan(&detail.OrderID, &detail.Price, &detail.Paid)
	if err != nil {
		return nil, err
	}
	return &detail, nil
}

func UpdatePriceDetail(ctx context.Context, detail models.PriceDetail) (*models.PriceDetail, error) {
	query := `UPDATE price_detail SET price = $1, paid = $2 WHERE order_id = $3`
	_, err := initializers.Pool.Exec(ctx, query, detail.Price, detail.Paid, detail.OrderID)
	if err != nil {
		return nil, err
	}
	return &detail, nil
}

func DeletePriceDetail(ctx context.Context, orderID int) error {
	query := `DELETE FROM price_detail WHERE order_id = $1`
	_, err := initializers.Pool.Exec(ctx, query, orderID)
	return err
}

