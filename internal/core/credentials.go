package core

import "time"

type Credentials struct {
	ID        int64     `db:"column:id;primary_key"`
	Timestamp time.Time `db:"column:timestamp"`
	Login     string    `db:"column:login"`
	Password  string    `db:"column:password"`
}

func (Credentials) TableName() string {
	return "credentials"
}
