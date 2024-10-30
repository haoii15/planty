package db

import (
	"context"
	"fmt"
	"planty/models"

	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
			return "", err
	}
	return string(hashedPassword), nil
}


func CreateUser(firstName, lastName, username, password string) error {
	ctx := context.Background()

	hashedPassword, err := hashPassword(password)
	if err != nil {
			return fmt.Errorf("failed to hash password: %w", err)
	}

	// Creating a new user instance
	user := &models.User{
			FirstName: firstName,
			LastName:  lastName,
			Username:  username,
			Password:  hashedPassword,
	}

	// Inserting the user into the database
	_, err = db.NewInsert().Model(user).Exec(ctx)
	if err != nil {
			return fmt.Errorf("failed to insert user: %w", err)
	}

	fmt.Printf("User %s created successfully!\n", username)
	return nil
}