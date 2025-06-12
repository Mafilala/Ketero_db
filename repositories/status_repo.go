
package repositories

import (
    "context"
    "errors"
    "github.com/Mafilala/ketero/backend/initializers"
    "github.com/Mafilala/ketero/backend/models"
)

func CreateNewStatus(ctx context.Context, name string) (*models.Status, error) {
    var newStatus models.Status
    query := `INSERT INTO status(name) VALUES ($1) RETURNING id, name`
    err := initializers.Pool.QueryRow(ctx, query, name).Scan(&newStatus.ID, &newStatus.Name)
    if err != nil {
        return nil, errors.New("unable to create new status")
    }
    return &newStatus, nil
}

func DeleteStatus(ctx context.Context, id int) (int, error) {
    var deletedID int
    query := `DELETE FROM status WHERE id=$1 RETURNING id`
    err := initializers.Pool.QueryRow(ctx, query, id).Scan(&deletedID)
    if err != nil {
        return id, errors.New("unable to delete status")
    }
    return deletedID, nil
}

func FindStatusByID(ctx context.Context, id int) (*models.Status, error) {
    var status models.Status
    query := `SELECT * FROM status WHERE id = $1`
    err := initializers.Pool.QueryRow(ctx, query, id).Scan(&status.ID, &status.Name)
    if err != nil {
        return nil, errors.New("status not found")
    }
    return &status, nil
}

func FindAllStatuses(ctx context.Context) (*[]models.Status, error) {
    var statusList []models.Status
    query := `SELECT * FROM status`
    rows, err := initializers.Pool.Query(ctx, query)
    if err != nil {
        return nil, errors.New("statuses not found")
    }

    for rows.Next() {
        var status models.Status
        err := rows.Scan(&status.ID, &status.Name)
        if err != nil {
            return nil, err
        }
        statusList = append(statusList, status)
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

    return &statusList, nil
}

func UpdateStatus(ctx context.Context, id int,  name string) (*models.Status, error) {
	var updatedStatus models.Status
	query := `UPDATE status SET name = $1 WHERE id = $2 RETURNING id, name`
	err := initializers.Pool.QueryRow(ctx, query, name, id).Scan(
		&updatedStatus.ID,
		&updatedStatus.Name,
	)
	if err != nil {
		return nil, errors.New("unable to update status")
	}
	return &updatedStatus, nil
}


