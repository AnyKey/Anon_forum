package model

type Message struct {
	ID         int64  `db:"id" json:"id"`
	Text       string `db:"text" json:"text"`
	ParentID   *int64 `db:"parent_id" json:"parent_id"`
	Rating     int    `db:"rating" json:"rating"`
	CategoryID int    `db:"category_id" json:"category_id"`
	ChildCount int    `db:"children_count" json:"children_count"`
}
