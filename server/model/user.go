package model

import (
	"github.com/jinzhu/gorm"
)

//https://github.com/rr-/szurubooru/blob/master/doc/API.md#user
type User struct {
	Version           int
	Name              string
	Password          string // not included in resource definition
	Rank              string
	LastLoginTime     time.Time
	CreationTime      time.Time
	AvartarStyle      string
	AvartarUrl        string
	CommentCount      int
	UploadedPostCount int
	LikedPostCount    int
	DislikedPostCount int
	FavoritePostCount int
	gorm.Model
}

func (User) TableName() string {
	return "user"
}

func (u *User) Get() error {
	return &res, db.Where("name = ?", u.Name).First(u).Error
}

func (User) GetAll() (*[]User, error) {
	res := make([]*User, 0)
	return res, db.Find(&res).Error
}

func (u *User) Create() error {
	return db.Create(u).Error
}

func (u *User) Update() error {
	return db.Model(u).Updates(u).Error
}

func (u *User) Delete(n) error {
	return db.Where("name = ?", u.Name).Delete(&User{}).Error
}
