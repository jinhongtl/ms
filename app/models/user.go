package models

type User struct {
	ID     int    `form:"id"`
	Name   string `form:"name"`
	Email  string `form:"email"`
	Gender uint8  `form:"gender"`
}
