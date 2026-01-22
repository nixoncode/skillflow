package courses

import "time"

type Course struct {
	ID          int64     `json:"id"`
	UserID      int64     `json:"user_id" db:"user_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Thumbnail   *string   `json:"thumbnail"` // pointer for nullable field
	Price       float64   `json:"price"`
	IsPublished bool      `json:"is_published" db:"is_published"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

type createCourseRequest struct {
	UserID      int64   `json:"user_id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Thumbnail   *string `json:"thumbnail"`
	Price       float64 `json:"price"`
	IsPublished bool    `json:"is_published"`
}
