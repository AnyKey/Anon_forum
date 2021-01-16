package handlers

import (
	"awesomeProject/model"
	"awesomeProject/repository"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type CategoryHandler struct {
	Conn *sqlx.DB
	Logger *zap.SugaredLogger
	Repo repository.CategoryRepository
}

func (handler CategoryHandler) GetCategories() ([]model.Category, error) {
	return handler.Repo.GetCategories()
}