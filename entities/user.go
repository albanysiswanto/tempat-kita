package entities

import "time"

type User struct {
	Id         uint
	Email      string
	Name       string
	Password   string
	Role       int
	Created_At time.Time
	Updated_At time.Time
}
