
package repositories

import (
	"context"
	"errors"

	"github.com/Mafilala/ketero/backend/initializers"
	"github.com/Mafilala/ketero/backend/models"
)

func AddMeasure(ctx context.Context, clothing_id, measure_id int) (*models.ClothingMeasures, error) {
	var newClothingMeasure models.ClothingMeasures
	query := `INSERT INTO clothing_measures(clothing_id, measure_id) VALUES ($1, $2) RETURNING clothing_id, measure_id`
	err := initializers.Pool.QueryRow(ctx, query, clothing_id, measure_id).Scan(&newClothingMeasure.Clothing_id, &newClothingMeasure.Measure_id)
	if err != nil {
		return nil, err 
	}
	return &newClothingMeasure, nil
}

func RemoveMeasure(ctx context.Context, clothing_id, measure_id int) (int, error) {
	var deletedID int
	query := `DELETE FROM clothing_measures WHERE clothing_id=$1 AND measure_id=$2 RETURNING measure_id`
	err := initializers.Pool.QueryRow(ctx, query, clothing_id, measure_id).Scan(&deletedID)
	if err != nil {
		return measure_id, errors.New("unable to delete measure")
	}
	return deletedID, nil
}

func GetAllClothingMeasures(ctx context.Context, clothing_id int) (*[]models.Measure, error) {
	allMeasures := []models.Measure{}

	query := `SELECT t2.* FROM clothing_measures t1 JOIN measure t2 ON t2.id = t1.measure_id WHERE t1.clothing_id = $1`
	rows, err := initializers.Pool.Query(ctx, query, clothing_id)
	if err != nil {
		return nil, errors.New("Unable to fetch measures")
	}

	for rows.Next() {
		var measure models.Measure
		err := rows.Scan(&measure.ID, &measure.Name)
		if err != nil {
			return nil, err
		}
		allMeasures = append(allMeasures, measure)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &allMeasures, nil
} 
