
package repositories

import (
	"context"
	"github.com/Mafilala/ketero/backend/initializers"
	"github.com/Mafilala/ketero/backend/models"
)

func CreateOrderDetail(ctx context.Context, detail models.OrderDetail) (*models.OrderDetail, error) {
	query := `INSERT INTO order_detail (order_id, style, fabric, color)
	          VALUES ($1, $2, $3, $4)`
	_, err := initializers.Pool.Exec(ctx, query,
		detail.OrderID, *detail.Style, *detail.Fabric, *detail.Color)
	if err != nil {
		return nil, err
	}
	return &detail, nil
}

func UpdateOrderDetail(ctx context.Context, detail models.OrderDetail) error {
	query := `UPDATE order_detail SET style = $1, fabric = $2, color = $3 WHERE order_id = $4`
	_, err := initializers.Pool.Exec(ctx, query,
		detail.Style, detail.Fabric, detail.Color, detail.OrderID)
	return err
}

func DeleteOrderDetail(ctx context.Context, orderID int) error {
	query := `DELETE FROM order_detail WHERE order_id = $1`
	_, err := initializers.Pool.Exec(ctx, query, orderID)
	return err
}

func GetOrderDetail(ctx context.Context, orderID int) (*models.OrderDetail, error) {
	query := `SELECT order_id, style, fabric, color FROM order_detail WHERE order_id = $1`
	row := initializers.Pool.QueryRow(ctx, query, orderID)

	var detail models.OrderDetail
	err := row.Scan(&detail.OrderID, &detail.Style, &detail.Fabric, &detail.Color)
	if err != nil {
		return nil, err
	}
	return &detail, nil
}
