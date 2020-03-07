package entity

import (
	"time"
)

type User struct {
	ID       uint
	FullName string `gorm:"type:varchar(255);not null;unique"`
	// UserName      string `gorm:"type:varchar(255);not null"`
	Password    string `gorm:"type:varchar(255);"`
	Email       string `gorm:"type:varchar(255);"`
	Phone       string `gorm:"type:varchar(255);"`
	Address     string `gorm:"type:varchar(255);"`
	Image       string `gorm:"type:varchar(255);"`
	Sex         string `gorm:"type:varchar(255);"`
	RoleId      uint
	BirthDate   time.Time
	Bio string
}
type Post struct{
	Uuid       uint
	Title string
	Body string
	Rate int
}
type Message struct{
	SenderId uint
	ReciverId uint
	Text string
	Date time.Time
}
type Role struct {
	ID    uint
	Name  string `gorm:"type:varchar(255)"`
	Users []User
}
type Session struct {
	ID         uint
	UUID       string `gorm:"type:varchar(255);not null"`
	Expires    int64  `gorm:"type:varchar(255);not null"`
	SigningKey []byte `gorm:"type:varchar(255);not null"`
}
