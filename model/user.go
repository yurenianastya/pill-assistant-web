package model

import (
	"awesomeProject/db"
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Account struct {
	bun.BaseModel `bun:"table:account"`

	ID       uuid.UUID `bun:",pk,type:uuid,default:uuid_generate_v4()"`
	Email    string    `bun:"email,notnull"`
	Password string    `bun:"pass,notnull"`
	Username string    `bun:"username,notnull"`
}

func GetAccounts() ([]Account, error) {
	var accounts []Account
	con, err := db.Connection()
	err = con.NewSelect().Model(&accounts).Scan(context.Background())
	if err != nil {
		return accounts, fmt.Errorf("error while selecting accounts from db")
	}
	err = con.Close()
	if err != nil {
		return accounts, fmt.Errorf("error while closing connection db")
	}
	return accounts, nil
}

func GetAccount(id uuid.UUID) (Account, error) {
	var account Account
	con, err := db.Connection()
	if err != nil {
		return account, fmt.Errorf("error while connecting to db")
	}
	err = con.NewSelect().Model(&account).
		Where("id = ?", id).Scan(context.Background())
	if err != nil {
		return account, fmt.Errorf("error while selecting account from db")
	}
	err = con.Close()
	if err != nil {
		return account, fmt.Errorf("error while closing connection db")
	}
	return account, nil
}

func CreateAccount(acc Account) error {
	con, err := db.Connection()
	if err != nil {
		return fmt.Errorf("error while connecting to db")
	}
	_, err = con.NewInsert().Model(&acc).Exec(context.Background())
	if err != nil {
		return fmt.Errorf("error while inserting account to db")
	}
	err = con.Close()
	if err != nil {
		return fmt.Errorf("error while closing connection db")
	}
	return nil
}

func DeleteAccount(id uuid.UUID) error {
	con, err := db.Connection()
	if err != nil {
		return fmt.Errorf("error while connecting to db")
	}
	_, err = con.NewDelete().Table("account").Where("id = ?", id).
		Exec(context.Background())
	if err != nil {
		return fmt.Errorf("error while deleting accoutn from db")
	}
	err = con.Close()
	if err != nil {
		return fmt.Errorf("error while closing connection db")
	}
	return nil
}
