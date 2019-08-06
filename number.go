package main

import "time"
import "github.com/google/uuid"

// Number represents number used by users to identify builds and other stuff
type Number struct {
	ID        string     `json:"id" gorm:"PRIMARY_KEY"`
	Number    int64      `json:"number"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

// NewNumber created a new instance of Number and saves it into the database
func NewNumber() (*Number, error) {
	number := Number{
		ID:        uuid.New().String(),
		Number:    0,
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	}

	err := db.Create(&number).Error

	return &number, err
}

// Find returns a single instance
func Find(ID string) (*Number, error) {
	var number = Number{}

	err := db.Where(&Number{ID: ID}).First(&number).Error
	if err != nil {
		return &number, err
	}

	return &number, nil
}

// Get returns value of the saved number
func (n *Number) Get() int64 {
	return n.Number
}

// Set sets new value of the number
func (n *Number) Set(newNumber int64) error {
	n.Number = newNumber
	return n.Save()
}

// Incr increases value of the number by 1
func (n *Number) Incr() error {
	n.Number++
	return n.Save()
}

// Decr decreses value of the number by 1
func (n *Number) Decr() error {
	n.Number--
	return n.Save()
}

// Save current value of the number
func (n *Number) Save() error {
	err := db.Save(n).Error
	return err
}
