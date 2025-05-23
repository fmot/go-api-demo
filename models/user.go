package models

import (
	"errors"

	"example.com/go-api-demo/db"
	"example.com/go-api-demo/models/utils"
)

type User struct {
  ID		int64
  Email		string `binding:"required"`
  Password 	string `binding:"required"`
}

func (u User) Save() error {
	query := "INSERT INTO users(email, password) VALUES (?, ?)"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashedPassword)

	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()

	u.ID = userId
	return err
}

func (u User) ValidateCredentials() error{
	query := "SELECT password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)

	var retrivePassword string
	err := row.Scan(&retrivePassword)

	if err != nil {
		return errors.New("Credentials invalid")
	}

	passwordIsValid := utils.CheckPasswordHash(u.Password, retrivePassword)
	
	if !passwordIsValid {
		return errors.New("Credentials invalid")
	}

	return nil
}