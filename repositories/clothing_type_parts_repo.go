
package repositories

import (
	"context"
	"errors"

	"github.com/Mafilala/ketero/backend/initializers"
	"github.com/Mafilala/ketero/backend/models"
)

func AddClothing(ctx context.Context, clothing_type_id, clothing_id int) (*models.ClothingTypePart, error) {
	var newClothingTypePart models.ClothingTypePart
	query := `INSERT INTO clothing_type_parts(clothing_type_id, clothing_id) VALUES ($1, $2) RETURNING clothing_type_id, clothing_id`
	err := initializers.Pool.QueryRow(ctx, query, clothing_type_id, clothing_id).Scan(&newClothingTypePart.Clothing_type_id, &newClothingTypePart.Clothing_id)
	if err != nil {
		return nil, err 
	}
	return &newClothingTypePart, nil
}

func RemoveClothingTypePart(ctx context.Context, clothing_type_id, clothing_id int) (int, error) {
	var deletedID int
	query := `DELETE FROM clothing_type_parts WHERE clothing_type_id=$1 AND clothing_id=$2 RETURNING clothing_id`
	err := initializers.Pool.QueryRow(ctx, query, clothing_type_id, clothing_id).Scan(&deletedID)
	if err != nil {
		return deletedID, errors.New("unable to delete client")
	}
	return deletedID, nil
}

func GetAllClothingPart(ctx context.Context, clothing_type_id int) ([]models.Clothing, error) {
    clothingParts := []models.Clothing{} // Initialize as empty slice
    query := `SELECT t2.* FROM clothing_type_parts t1 JOIN clothing t2 ON t2.id = t1.clothing_id WHERE clothing_type_id = $1`
    
    rows, err := initializers.Pool.Query(ctx, query, clothing_type_id)
    if err != nil {
        return nil, errors.New("Unable to fetch clothing parts")
    }
    defer rows.Close() // Important: close rows to prevent connection leaks

    for rows.Next() {
        var clothingPart models.Clothing
        if err := rows.Scan(&clothingPart.ID, &clothingPart.Name); err != nil {
            return nil, err
        }
        clothingParts = append(clothingParts, clothingPart)
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

    return clothingParts, nil
} 
