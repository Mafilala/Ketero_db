package repositories

import (
	"context"
	"errors"

	"github.com/Mafilala/ketero/backend/initializers"
	"github.com/Mafilala/ketero/backend/models"
)

func CreateNewUser(ctx context.Context, user *models.User) (*models.User, error) {
	var newUser models.User
	query := `INSERT INTO users(telegram_id, name, role) VALUES ($1, $2, $3) RETURNING id, telegram_id, name, role`
	err := initializers.Pool.QueryRow(ctx, query, user.TelegramID, user.Name, user.Role).Scan(&newUser.TelegramID, &newUser.Name, &newUser.Role)
	if err != nil {
		return nil, err
	}
	return &newUser, nil
}

func DeleteUser(ctx context.Context, id int) (int, error) {
	var deletedID int
	query := `DELETE FROM users WHERE id=$1 RETURNING id`
	err := initializers.Pool.QueryRow(ctx, query, id).Scan(&deletedID)
	if err != nil {
		return id, errors.New("unable to delete user")
	}
	return deletedID, nil
}

func FindUserByID(ctx context.Context, id int) (*models.User, error) {
	var user models.User
	query := `SELECT * FROM users WHERE telegram_id = $1`
	err := initializers.Pool.QueryRow(ctx, query, id).Scan(&user.ID, &user.TelegramID, &user.Name, &user.Role)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return &user, nil
}

func GetAllUser(ctx context.Context) (*[]models.User, error) {
	var userList []models.User
	query := `SELECT * FROM users`
	rows, err := initializers.Pool.Query(ctx, query)
	if err != nil {
		return nil, errors.New("users not found")
	}

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.TelegramID, &user.Name, &user.Role)
		if err != nil {
			return nil, err
		}
		userList = append(userList, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &userList, nil
}
