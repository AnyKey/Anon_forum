package model

type Category struct {
	ID          int    `db:"id" json:"id"`
	Title       string `db:"title" json:"title"`
	Description *string `db:"description" json:"description"`
}

type AddMessagePostRequest struct {
	CategoryID int    `json:"category_id"`
	Text       string `json:"text"`
	ParentID   *int64 `json:"parent_id"`
}
