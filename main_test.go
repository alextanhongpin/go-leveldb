package main_test

import (
	"log"
	"testing"

	"github.com/stretchr/testify/mock"
)

type DB struct {
	mock.Mock
}

func (d *DB) Has(id string, in int) (bool, int) {
	args := d.Called(id, in)
	return args.Bool(0), args.Int(1)
}

func TestDB(t *testing.T) {

	db := new(DB)
	db.On("Has", "0", 0).Return(true, 100)
	db.On("Has", "0", 1).Return(false, 200)
	db.On("Has", "1", 2).Return(false, 300)

	ok, val := db.Has("0", 0)
	log.Println(ok, val)

	ok, val = db.Has("0", 1)
	log.Println(ok, val)
}
