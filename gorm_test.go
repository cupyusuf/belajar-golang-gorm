package belajar_golang_gorm

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func OpenConnection() *gorm.DB {
	dialect := mysql.Open("root:@tcp(localhost:3306)/belajar_golang_gorm?charset=utf8mb4&parseTime=True&loc=Local")
	db, err := gorm.Open(dialect, &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}

var db = OpenConnection()

func TestOpenConnection(t *testing.T) {
	assert.NotNil(t, db)
}

func TestExecuteSQL(t *testing.T) {
	err := db.Exec("insert into sample(id, name) values (?, ?)", "1", "Yusuf").Error
	assert.Nil(t, err)

	err = db.Exec("insert into sample(id, name) values (?, ?)", "2", "Firlana").Error
	assert.Nil(t, err)

	err = db.Exec("insert into sample(id, name) values (?, ?)", "3", "Supriadi").Error
	assert.Nil(t, err)

	err = db.Exec("insert into sample(id, name) values (?, ?)", "4", "Luchiana").Error
	assert.Nil(t, err)
}
