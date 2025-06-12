
package repositories

import (
	"context"
	"errors"
        "github.com/Mafilala/ketero/backend/initializers"
	"github.com/Mafilala/ketero/backend/models"
)


func CreateNewClothingType(ctx context.Context, name string) (*models.ClothingType, error) {
	var newClothingType models.ClothingType
	query := `INSERT INTO  clothing_type(name) VALUES ($1) RETURNING id, name` 
	err := initializers.Pool.QueryRow(ctx, query, name).Scan(&newClothingType.ID, &newClothingType.Name)
	if err != nil {
		return nil, errors.New("unable to create new clothing type") 
	}
	
	return &newClothingType, nil

}

func UpdateClothingType(ctx context.Context, id int,  name string) (*models.ClothingType, error) {
	var newClothingType models.ClothingType
	query := `UPDATE clothing_type SET name = $1 WHERE id = $2 RETURNING id, name`
	err := initializers.Pool.QueryRow(ctx, query, name, id).Scan(
		&newClothingType.ID,
		&newClothingType.Name,
	)
	if err != nil {
		return nil, errors.New("unable to create new clothing type item")
	}
	return &newClothingType, nil
}



func DeleteClothingType(ctx context.Context, id int) (int, error) {
	var deleted_id int
	query := `DELETE FROM clothing_type WHERE id=$1 RETURNING id` 
	err := initializers.Pool.QueryRow(ctx, query, id).Scan(&deleted_id)
	if err != nil {
		return id, errors.New("unable to delete clothing type") 
	}
	
	return deleted_id, nil

}

func FindClothingTypeByID(ctx context.Context, id int) (*models.ClothingType, error) {
	var clothing_type models.ClothingType

	query := `SELECT * FROM clothing_type WHERE id = $1 `	
        err := initializers.Pool.QueryRow(ctx, query, id).Scan(&clothing_type.ID, &clothing_type.Name)
	
        if err != nil {
		return nil, errors.New("measure not found")
	}

	return &clothing_type, nil
}

func FindAllClothingType(ctx context.Context) (*[]models.ClothingType, error) {
	var clothing_type_list []models.ClothingType

	query := `SELECT * FROM clothing_type`	
        rows, err := initializers.Pool.Query(ctx, query)


        if err != nil {
		return nil, errors.New("clothing types not found")
	}

	for rows.Next() {
		var clothing_type models.ClothingType
		err := rows.Scan(&clothing_type.ID, &clothing_type.Name)
		
		if err != nil {
			return nil, err
		}

		clothing_type_list = append(clothing_type_list, clothing_type)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &clothing_type_list, nil
}




