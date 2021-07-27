package input

import "time"

type User struct {
	Id        int        `json:"id"`
	UserName  string     `json:"user_name"`
	Birthday  time.Time  `json:"birthday"`
	Gender    uint8      `json:"gender"`
	Password  string     `json:"password"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
