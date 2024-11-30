package controllers

import (
	"go-crud-1/config"
	"go-crud-1/models"
)

func FetchUsers() ([]models.User, error) {
	rows, err := config.DB.Query("SELECT id, name, email, gender FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Gender); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
