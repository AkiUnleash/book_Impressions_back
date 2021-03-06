package models

import "time"

type Impression struct {
	Id        uint       `json:"id" param:"id"`
	Uid       string     `json:"uid"`
	Bookid    string     `json:"bookid"`
	Isbn10    string     `json:"isbn10"`
	Isbn13    string     `json:"isbn13"`
	Booktitle string     `json:"booktitle"`
	Imageurl  string     `json:"imageurl" gorm:"size:2083"`
	Title     string     `json:"title"`
	Body      string     `json:"body" gorm:"type:mediumtext"`
	CreatedAt time.Time  `json:"create_at"`
	UpdatedAt time.Time  `json:"update_at"`
	DeletedAt *time.Time `json:"delete_at"`
}
