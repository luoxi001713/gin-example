package models

import (
	"errors"
	"time"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID               int64       `gorm:"primaryKey" db:"id" json:"id"`
	Name             string      `db:"name" json:"name"`
	CreateTime       time.Time   `db:"create_time" json:"create_time"`
}

func (b *User) TableName() string {
	return "example_db.user"
}
