
package services

import (
	"context"

	"github.com/Mafilala/ketero/backend/models"
	"github.com/Mafilala/ketero/backend/repositories"
)

func GetClientByID(ctx context.Context, id int) (*models.Client, error) {
	return repositories.FindClientByID(ctx, id)
}

func CreateNewClient(ctx context.Context, name, phone string) (*models.Client, error) {
	return repositories.CreateNewClient(ctx, name, phone)
}

func DeleteClient(ctx context.Context, id int) (int, error) {
	return repositories.DeleteClient(ctx, id)
}

func GetAllClients(ctx context.Context) (*[]models.Client, error) {
	return repositories.FindAllClients(ctx)
}
