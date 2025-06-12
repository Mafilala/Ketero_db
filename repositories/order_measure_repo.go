
package repositories

import (
	"context"
	"github.com/Mafilala/ketero/backend/initializers"
	"github.com/Mafilala/ketero/backend/models"
)

func CreateOrderMeasure(ctx context.Context, om models.OrderMeasure) (*models.OrderMeasure, error) {
	query := `INSERT INTO order_measure (order_id, measure_id, measure, clothing_id)
			  VALUES ($1, $2, $3, $4)`
	_, err := initializers.Pool.Exec(ctx, query, om.OrderID, om.MeasureID, om.Measure, om.ClothingID)
	if err != nil {
		return nil, err
	}
	return &om, nil
}

func UpdateOrderMeasure(ctx context.Context, orderID, clothingID int, measures []models.MeasureValue) error {
	query := `UPDATE order_measure SET measure = $1 WHERE order_id = $2 AND measure_id = $3 AND clothing_id = $4`

	for _, m := range measures {
		_, err := initializers.Pool.Exec(ctx, query, m.Value, orderID, m.MeasureID, clothingID)
		if err != nil {
			return err
		}
	}

	return nil
}
func DeleteOrderMeasure(ctx context.Context, orderID, measureID int) error {
	query := `DELETE FROM order_measure WHERE order_id = $1 AND measure_id = $2`
	_, err := initializers.Pool.Exec(ctx, query, orderID, measureID)
	return err
}

func GetOrderMeasuresByOrderID(ctx context.Context, orderID int) ([]models.OrderMeasure, error) {
	query := `SELECT order_id, measure_id, measure, clothing_id FROM order_measure WHERE order_id = $1`
	rows, err := initializers.Pool.Query(ctx, query, orderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var measures []models.OrderMeasure
	for rows.Next() {
		var om models.OrderMeasure
		if err := rows.Scan(&om.OrderID, &om.MeasureID, &om.Measure, &om.ClothingID); err != nil {
			return nil, err
		}
		measures = append(measures, om)
	}
	return measures, nil
}
