package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	//	_ "github.com/jinzhu/gorm/dialects/mssql"
	//	_ "github.com/jinzhu/gorm/dialects/mysql"
	//	_ "github.com/jinzhu/gorm/dialects/postgres"
	"fmt"
)

type DbKind int

const (
	DB_KIND_SQLITE DbKind = iota
	DB_KIND_MYSQL
)

type DbOption struct {
	Kind     DbKind // adapter type
	Name     string // database name
	UserName string // username of connection
	UserPass string // password of connection
}

var db *gorm.DB

// Open gorm object and store it in internal variable
func Open(opt DbOption) error {
	var dia, conn string

	switch opt.Kind {
	case DB_KIND_SQLITE:
		dia = "sqlite3"
		conn = opt.Name
	case DB_KIND_MYSQL:
		dia = "mysql"
		conn = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", opt.UserName, opt.UserPass, opt.Name)
	}

	d, err := gorm.Open(dia, conn)
	if err != nil {
		return err
	} else {
		db = d
		db.AutoMigrate(&TagCategory{})
		return nil
	}
}

func Close() error {
	return db.Close()
}
