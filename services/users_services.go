package services

import (
	"context"

	"github.com/Mafilala/ketero/backend/models"
	"github.com/Mafilala/ketero/backend/repositories"
)

func GetUserByID(ctx context.Context, id int) (*models.User, error) {
	return repositories.FindUserByID(ctx, id)
}

func CreateNewUser(ctx context.Context, user *models.User) (*models.User, error) {
	return repositories.CreateNewUser(ctx, user)
}

func DeleteUser(ctx context.Context, id int) (int, error) {
	return repositories.DeleteUser(ctx, id)
}
