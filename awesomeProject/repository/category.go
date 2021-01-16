package repository

import (
	"awesomeProject/model"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type CategoryRepository struct {
	Conn *sqlx.DB
}

func (repo CategoryRepository) GetCategories() ([]model.Category, error) {
	var thread []model.Category
	err := repo.Conn.Select(&thread, "select * from category")
	if err != nil {
		return nil, errors.Wrap(err, "error select categories")
	}

	return thread, nil
}
