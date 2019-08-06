package main

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jinzhu/gorm"
)

func TestMain(m *testing.M) {
	var err error

	// Prepare database connection
	db, err = gorm.Open("sqlite3", "gorm.db")
	if err != nil {
		log.Fatalln("Database connection issue:", err)
	}
	defer db.Close()

	db.AutoMigrate(&Number{})

	// Run the test
	code := m.Run()
	os.Exit(code)
}

func TestNumber(t *testing.T) {
	number, err := NewNumber()
	assert.Nil(t, err)
	assert.Equal(t, int64(0), number.Number)

	err = number.Incr()
	assert.Nil(t, err)
	assert.Equal(t, int64(1), number.Number)

	err = number.Set(99)
	assert.Nil(t, err)
	assert.Equal(t, int64(99), number.Number)

	err = number.Decr()
	assert.Nil(t, err)
	assert.Equal(t, int64(98), number.Number)

	value := number.Get()
	assert.Nil(t, err)
	assert.Equal(t, int64(98), value)

	number, err = Find(number.ID)
	assert.Nil(t, err)
	assert.Equal(t, int64(98), value)
}
