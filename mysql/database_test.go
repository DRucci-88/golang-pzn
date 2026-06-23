package main

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestOpenConnection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

}
