package controller

import (
	"awesomeProject/db"
	"awesomeProject/model"
	"context"
	"fmt"
)

func FindByField(fieldName, field string) (string, error) {
	acc := *new(model.Account)
	ctx := context.Background()
	con, err := db.Connection()
	if err != nil {
		return "", fmt.Errorf("error while connecting to db")
	}
	err = con.NewSelect().Model(&acc).Where(fieldName+" = ?", field).Scan(ctx)
	if err != nil {
		return "", fmt.Errorf("error while selecting username from db")
	}
	err = con.Close()
	if err != nil {
		return "", fmt.Errorf("error while closing connection db")
	}
	return field, nil
}

func SaveUser(username, password, email string) error {
	acc := model.Account{
		Username: username,
		Password: password,
		Email:    email,
	}
	err := model.CreateAccount(acc)
	if err != nil {
		return fmt.Errorf("error while writing user to db")
	}
	return nil
}

// TODO write funcs for change username, change password
