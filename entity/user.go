package entity

import "time"

type UserEntity struct {
	ID        string    `gorm:"primary_key"`
	Email     string    `gorm:"column:email;unique"`
	Password  string    `gorm:"column:password"`
	Name      string    `gorm:"column:name"`
	Age       int8      `gorm:"column:age"`
	Status    bool      `gorm:"column:status"`
	CreatedAt time.Time `gorm:"column:created_at"`
}
