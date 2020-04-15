package model

import (
	"encoding/json"
	//"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"net/http"
	"reflect"
)

/*
Dynamic type assertion is impossible in golang.
So, need to write type switch per every type or waitng generic update.
*/

func ListingModel(model interface{}) func(echo.Context) error {
	return func(c echo.Context) error {
		res := newByType(model)
		db.Find(&res)
		return c.JSON(http.StatusOK, res)
	}
}

func CreateModel(model interface{}) func(echo.Context) error {
	return func(c echo.Context) error {
		m := newByType(model)
		if err := json.NewDecoder(c.Request().Body).Decode(&m); err != nil {
			return err
		}
		db.Create(m)
		return c.JSON(http.StatusOK, m)
	}
}

func newByType(object interface{}) interface{} {
	e := reflect.TypeOf(object)
	return reflect.New(e).Elem().Interface()
}

func getType(T interface{}) string {
	if t := reflect.TypeOf(T); t.Kind() == reflect.Ptr {
		return "*" + t.Elem().Name()
	} else {
		return t.Name()
	}
}
