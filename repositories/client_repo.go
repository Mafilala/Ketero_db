
package repositories

import (
	"context"
	"errors"

	"github.com/Mafilala/ketero/backend/initializers"
	"github.com/Mafilala/ketero/backend/models"
)

func CreateNewClient(ctx context.Context, name, phone string) (*models.Client, error) {
	var newClient models.Client
	query := `INSERT INTO client(full_name, phone_number) VALUES ($1, $2) RETURNING id, full_name, phone_number`
	err := initializers.Pool.QueryRow(ctx, query, name, phone).Scan(&newClient.ID, &newClient.Name, &newClient.Phone)
	if err != nil {
		return nil, err 
	}
	return &newClient, nil
}

func DeleteClient(ctx context.Context, id int) (int, error) {
	var deletedID int
	query := `DELETE FROM client WHERE id=$1 RETURNING id`
	err := initializers.Pool.QueryRow(ctx, query, id).Scan(&deletedID)
	if err != nil {
		return id, errors.New("unable to delete client")
	}
	return deletedID, nil
}

func FindClientByID(ctx context.Context, id int) (*models.Client, error) {
	var client models.Client
	query := `SELECT * FROM client WHERE id = $1`
	err := initializers.Pool.QueryRow(ctx, query, id).Scan(&client.ID, &client.Name, &client.Phone)
	if err != nil {
		return nil, errors.New("client not found")
	}
	return &client, nil
}

func FindAllClients(ctx context.Context) (*[]models.Client, error) {
	var clientList []models.Client
	query := `SELECT * FROM client`
	rows, err := initializers.Pool.Query(ctx, query)
	if err != nil {
		return nil, errors.New("clients not found")
	}

	for rows.Next() {
		var client models.Client
		err := rows.Scan(&client.ID, &client.Name, &client.Phone)
		if err != nil {
			return nil, err
		}
		clientList = append(clientList, client)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &clientList, nil
}
