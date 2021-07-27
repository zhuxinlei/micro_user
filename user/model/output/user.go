package output

import "time"

type User struct {
	Id        int       `json:"id"`
	UserName  string    `json:"user_name"`
	Birthday  time.Time `json:"birthday"`
	Gender    uint8     `json:"gender"`
}
