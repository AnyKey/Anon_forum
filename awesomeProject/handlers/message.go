package handlers

import (
	"awesomeProject/model"
	"awesomeProject/repository"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type MessageHandler struct {
	Conn *sqlx.DB
	Logger *zap.SugaredLogger
	Repo repository.MessageRepository
}

func (handler MessageHandler) GetThreads(categoryID int) ([]model.Message, error) {
	return handler.Repo.GetThreadMessages(categoryID)
}

func (handler MessageHandler) GetMessagesByThread(categoryID int,parentID int64) ([]model.Message, error) {
	return handler.Repo.GetMessages(categoryID, &parentID)
}
func (handler MessageHandler) AddMessage(categoryID int, parentID *int64, text string) error {
	return handler.Repo.AddMessage(categoryID,parentID,text)
}