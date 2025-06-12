
package repositories

import (
	"context"
	"errors"
        "github.com/Mafilala/ketero/backend/initializers"
	"github.com/Mafilala/ketero/backend/models"
)

func CreateNewMeasure(ctx context.Context, name string) (*models.Measure, error) {
	var new_measure models.Measure
	query := `INSERT INTO  measure(name) VALUES ($1) RETURNING id, name` 
	err := initializers.Pool.QueryRow(ctx, query, name).Scan(&new_measure.ID, &new_measure.Name)
	if err != nil {
		return nil, errors.New("unable to create new measure") 
	}
	
	return &new_measure, nil

}

func FindMeasureByID(ctx context.Context, id int) (*models.Measure, error) {
	var measure models.Measure

	query := `SELECT * FROM measure WHERE id = $1 `	
        err := initializers.Pool.QueryRow(ctx, query, id).Scan(&measure.ID, &measure.Name)
	
        if err != nil {
		return nil, errors.New("measure not found")
	}

	return &measure, nil
}

func FindAllMeasure(ctx context.Context) (*[]models.Measure, error) {
	var measures []models.Measure

	query := `SELECT * FROM measure`	
        rows, err := initializers.Pool.Query(ctx, query)


        if err != nil {
		return nil, errors.New("measure not found")
	}

	for rows.Next() {
		var measure models.Measure
		err := rows.Scan(&measure.ID, &measure.Name)
		
		if err != nil {
			return nil, err
		}

		measures = append(measures, measure)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &measures, nil
}

func DeleteMeasure(ctx context.Context, id int) (int, error) {
	var measure_id int
	query := `DELETE FROM measure WHERE id=$1 RETURNING id`

	err := initializers.Pool.QueryRow(ctx, query, id).Scan(&measure_id)
	
	if err != nil {
		return id, errors.New("Unable to delete the measure")
	}
	return measure_id, nil 
}

func UpdateMeasure(ctx context.Context, id int,  name string) (*models.Measure, error) {
	var newMeasure models.Measure
	query := `UPDATE measure SET name = $1 WHERE id = $2 RETURNING id, name`
	err := initializers.Pool.QueryRow(ctx, query, name, id).Scan(
		&newMeasure.ID,
		&newMeasure.Name,
	)
	if err != nil {
		return nil, errors.New("unable to update measure")
	}
	return &newMeasure, nil
}
