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

func NewTC(name, color string) *TagCategory {
	return &TagCategory{0, name, color, 0, false, gorm.Model{}}
}

func GetAllTC() ([]*TagCategory, error) {
	res := make([]*TagCategory, 0)
	return res, db.Find(&res).Error
}

func GetTC(name string) (*TagCategory, error) {
	res := TagCategory{}
	return &res, db.Where("name = ?", name).First(&res).Error
}

func CreateTC(tc *TagCategory) error {
	return db.Create(tc).Error
}

func UpdateTC(name string, color string, version int) error {
	if tc, err := GetTC(name); err != nil {
		return err
	} else {
		tc.Color = color
		tc.Version = version
		return db.Save(&tc).Error
	}
}

func DeleteTC(name string) error {
	return db.Where("name = ?", name).Delete(&TagCategory{}).Error
}

func SetDefaultTC(name string) (*TagCategory, error) {
	if tc, err := GetTC(name); err != nil {
		return nil, err
	} else {
		tc.Default = true
		return tc, db.Save(&tc).Error
	}
}
