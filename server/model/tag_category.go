package model

import (
	"github.com/jinzhu/gorm"
)

type TagCategory struct {
	Version int
	Name    string `gorm:"PRIMARY_KET;UNIQUE"`
	Color   string
	Usages  int
	Default bool
	gorm.Model
}

func (TagCategory) TableName() string {
	return "table_category"
}

func (TagCategory) GetAll() ([]*TagCategory, error) {
	res := make([]*TagCategory, 0)
	return res, db.Find(&res).Error
}

func (tc *TagCategory) Get() (*TagCategory, error) {
	return &res, db.Where("name = ?", tc.Name).First(tc).Error
}

func (tc *TagCategory) Create() error {
	return db.Create(tc).Error
}

func (tc *TagCategory) Update() error {
	return db.Model(tc).Updates(tc).Error
}

func (tc *TagCategory) Delete() error {
	return db.Where("name = ?", tc.Name).Delete(tc).Error
}

func (TagCategory) SetDefault(name string) (*TagCategory, error) {
	if tc, err := GetTC(name); err != nil {
		return nil, err
	} else {
		tc.Default = true
		return tc, db.Save(&tc).Error
	}
}
