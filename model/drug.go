package model

import (
	"awesomeProject/db"
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Drug struct {
	bun.BaseModel `bun:"table:drug"`

	ID              uuid.UUID `bun:",pk,type:uuid,default:uuid_generate_v4()"`
	activeSubstance string    `bun:"substance,notnull"`
	brandName       string    `bun:"brand,notnull"`
	commonName      string    `bun:"common_name"`
	dosage          int       `bun:"dosage,notnull"`
}

func GetDrugs() ([]Drug, error) {
	var drugs []Drug
	con, err := db.Connection()
	if err != nil {
		return drugs, fmt.Errorf("error while connecting to db")
	}
	err = con.NewSelect().Model(&drugs).Scan(context.Background())
	if err != nil {
		return drugs, fmt.Errorf("error while selecting drugs from db")
	}
	err = con.Close()
	if err != nil {
		return drugs, fmt.Errorf("error while closing connection db")
	}
	return drugs, nil
}

func GetDrug(id int) (*Drug, error) {
	drug := new(Drug)
	con, err := db.Connection()
	if err != nil {
		return drug, fmt.Errorf("error while connecting to db")
	}
	err = con.NewSelect().Model(drug).Where("id = ?", id).
		Scan(context.Background())
	if err != nil {
		return drug, fmt.Errorf("error while selecting drug from db")
	}
	err = con.Close()
	if err != nil {
		return drug, fmt.Errorf("error while closing connection db")
	}
	return drug, nil
}

func CreateDrug(drug Drug) error {
	con, err := db.Connection()
	if err != nil {
		return fmt.Errorf("error while connecting to db")
	}
	_, err = con.NewInsert().Model(&drug).Exec(context.Background())
	if err != nil {
		fmt.Errorf("error while creating drug in db")
	}
	err = con.Close()
	if err != nil {
		return fmt.Errorf("error while closing connection db")
	}
	return nil
}

func DeleteDrug(id int) error {
	con, err := db.Connection()
	if err != nil {
		return fmt.Errorf("error while connecting to db")
	}
	_, err = con.NewDelete().Table("account").
		Where("id = ?", id).Exec(context.Background())
	if err != nil {
		return fmt.Errorf("error while deleting drug from db")
	}
	err = con.Close()
	if err != nil {
		return fmt.Errorf("error while closing connection db")
	}
	return nil
}
