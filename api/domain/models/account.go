package models

import "time"

type Account struct {
	Id        uint       `json:"id"`
	Uid       string     `json:"uid" gorm:"unique"`
	Username  string     `json:"username"`
	Email     string     `json:"email" gorm:"unique"`
	Password  []byte     `json:"-"`
	CreatedAt time.Time  `json:"create_at"`
	UpdatedAt time.Time  `json:"update_at"`
	DeletedAt *time.Time `json:"delete_at"`
}

type AccountUpdate struct {
	Username string `json:"username"`
	Email    string `json:"email" gorm:"unique"`
}
