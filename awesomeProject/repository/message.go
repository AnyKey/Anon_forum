package repository

import (
	"awesomeProject/model"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type MessageRepository struct {
	Conn *sqlx.DB
}

func (repo MessageRepository) AddMessage(categoryID int, parentID *int64, text string) error {
	var pID sql.NullInt64
	if parentID == nil {
		pID = sql.NullInt64{}
	} else {
		pID = sql.NullInt64{Int64: *parentID, Valid: true}
	}
	res, err := repo.Conn.Exec("insert into message (category_id, parent_id, text) values(? , ? , ?)", categoryID, pID, text)
	if err != nil {
		return errors.Wrap(err, "error insert message")
	}

	if res != nil {
		if count, err := res.RowsAffected(); err != nil || count == 0 {
			return errors.Wrap(err, "error insert message row affected=0")
		}
	}

	return nil
}

func (repo MessageRepository) GetThreadMessages(categoryID int) ([]model.Message, error) {
	var thread []model.Message
	err := repo.Conn.Select(&thread, "select id, text,parent_id, rating,category_id,(select count(id) from message where parent_id=msg1.id) as children_count from message as msg1 where category_id=? and parent_id is null", categoryID)
	if err != nil {
		return nil, errors.Wrap(err, "error select thread messages")
	}

	return thread, nil
}

func (repo MessageRepository) GetMessages(categoryID int, parentID *int64) ([]model.Message, error) {
	var thread []model.Message
	err := repo.Conn.Select(&thread, "select id, text,parent_id, rating,category_id,(select count(id) from message where parent_id=msg1.id) as children_count from message as msg1 where category_id=? and parent_id=?", categoryID, parentID)
	if err != nil {
		return nil, errors.Wrap(err, "error select thread messages")
	}

	return thread, nil
}
