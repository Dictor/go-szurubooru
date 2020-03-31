package model_test

import (
	"github.com/dictor/go-szurubooru/model"
	//"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

type BasicModel interface {
	Get() error
	Create() error
	Delete() error
	Update() error
	TableName() string
}

var TestSet [][]interface{} = [][]interface{}{
	model.NewTC("rating", "red"),
	model.NewTC("file_size", "blue"),
	model.NewTC("artist", "green"),
}
var dbDir string

func TestOpenDB(t *testing.T) {
	dbDir, err := ioutil.TempDir("", "test")
	assert.NoError(t, err, "Making temp directory fail")
	assert.NoError(t, model.Open(model.DbOption{model.DB_KIND_SQLITE, dbDir + "/test.db", "", ""}), "Open db fail")
}

func TestInsertTC(t *testing.T) {
	for _, v := range TCset {
		model.CreateTC(v)
	}
	res, err := model.GetAllTC()
	assert.NoError(t, err)
	assert.Len(t, res, len(TCset))
}

func TestDeleteTC(t *testing.T) {
	assert.NoError(t, model.DeleteTC("rating"))
	res, err := model.GetAllTC()
	assert.NoError(t, err)
	assert.Len(t, res, len(TCset)-1)
}

func TestUpdateTC(t *testing.T) {
	assert.NoError(t, model.UpdateTC("artist", "orange", 1))
	res, err := model.GetTC("artist")
	assert.NoError(t, err)
	assert.Equal(t, res.Color, "orange")
	assert.Equal(t, res.Version, 1)
}

func TestSetDefaultTC(t *testing.T) {
	res, err := model.SetDefaultTC("file_size")
	assert.NoError(t, err)
	assert.Equal(t, res.Default, true)
}

func TestCloseDB(t *testing.T) {
	assert.NoError(t, model.Close())
	os.RemoveAll(dbDir)
}
