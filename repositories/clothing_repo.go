package repositories

import (
	"context"
	"errors"

	"github.com/Mafilala/ketero/backend/initializers"
	"github.com/Mafilala/ketero/backend/models"
)

func CreateNewClothing(ctx context.Context, name string) (*models.Clothing, error) {
	var newClothing models.Clothing
	query := `INSERT INTO clothing(name) VALUES ($1) RETURNING id, name`
	err := initializers.Pool.QueryRow(ctx, query, name).Scan(
		&newClothing.ID,
		&newClothing.Name,
	)
	if err != nil {
		return nil, errors.New("unable to create new clothing item")
	}
	return &newClothing, nil
}

func UpdateClothing(ctx context.Context, id int,  name string) (*models.Clothing, error) {
	var newClothing models.Clothing
	query := `UPDATE clothing SET name = $1 WHERE id = $2 RETURNING id, name`
	err := initializers.Pool.QueryRow(ctx, query, name, id).Scan(
		&newClothing.ID,
		&newClothing.Name,
	)
	if err != nil {
		return nil, errors.New("unable to create new clothing item")
	}
	return &newClothing, nil
}


func DeleteClothing(ctx context.Context, id int) (int, error) {
	var deletedID int
	query := `DELETE FROM clothing WHERE id = $1 RETURNING id`
	err := initializers.Pool.QueryRow(ctx, query, id).Scan(&deletedID)
	if err != nil {
		return id, errors.New("unable to delete clothing item")
	}
	return deletedID, nil
}

func FindClothingByID(ctx context.Context, id int) (*models.Clothing, error) {
	var clothing models.Clothing
	query := `SELECT id, name FROM clothing WHERE id = $1`
	err := initializers.Pool.QueryRow(ctx, query, id).Scan(
		&clothing.ID,
		&clothing.Name,
	)
	if err != nil {
		return nil, errors.New("clothing item not found")
	}
	return &clothing, nil
}

func FindAllClothing(ctx context.Context) (*[]models.Clothing, error) {
	var clothingList []models.Clothing
	query := `SELECT id, name FROM clothing`
	rows, err := initializers.Pool.Query(ctx, query)
	if err != nil {
		return nil, errors.New("failed to fetch clothing items")
	}
	defer rows.Close()

	for rows.Next() {
		var clothing models.Clothing
		err := rows.Scan(&clothing.ID, &clothing.Name)
		if err != nil {
			return nil, err
		}
		clothingList = append(clothingList, clothing)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &clothingList, nil
}

