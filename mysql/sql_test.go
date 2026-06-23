package main

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "INSERT INTO customer(id, name) VALUES('leas','LEAS')"

	_, err := db.ExecContext(ctx, script)
	if err != nil {
		panic(err)
	}
	fmt.Println("Success insert new customer")
}

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT id,name,email,balance,rating,birth_date,married,created_at FROM customer"

	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance int32
		var rating float64
		// var birthDate, createdAt time.Time
		var birthDate, createdAt sql.NullTime
		var married bool
		err := rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt)
		if err != nil {
			panic(err)
		}
		fmt.Println("id", id)
		fmt.Println("Name :", name)
		if email.Valid {
			fmt.Println("email :", email.String)
		}

		fmt.Println("balance :", balance)
		fmt.Println("rating :", rating)
		if birthDate.Valid {
			fmt.Println("birthDate", birthDate.Time)
		}

		fmt.Println("married :", married)
		fmt.Println("createdAt :", createdAt)
	}

	fmt.Println("Done Completed ")
}

func TestSqlInject(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin"
	password := "admin"

	/// SQL Injection
	username = "admin'; #"

	script := fmt.Sprintf("SELECT username, password FROM user WHERE username = '%s' AND password = '%s'", username, password)
	fmt.Println(script)
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username, password string
		err := rows.Scan(&username, &password)
		if err != nil {
			panic(err)
		}
		fmt.Println("username : ", username)
		fmt.Println("password :", password)
		fmt.Println("Sukses Login")
	} else {
		fmt.Println("Gagal Login")
	}

	fmt.Println("Done Completed")
}

func TestSqlInjectSafe(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin"
	password := "admin"

	/// SQL Injection
	username = "admin'; #"

	// script := fmt.Sprintf("SELECT username, password FROM user WHERE username = '%s' AND password = '%s'", username, password)
	script := "SELECT username, password FROM user WHERE username = ? AND password = ? LIMIT 1"

	fmt.Println(script)
	rows, err := db.QueryContext(ctx, script, username, password)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username, password string
		err := rows.Scan(&username, &password)
		if err != nil {
			panic(err)
		}
		fmt.Println("username : ", username)
		fmt.Println("password :", password)
		fmt.Println("Sukses Login")
	} else {
		fmt.Println("Gagal Login")
	}

	fmt.Println("Done Completed")
}

func TestAutoIncrement(t *testing.T) {
	db := GetConnection()

	defer db.Close()

	ctx := context.Background()

	email := "le2@gmail.com"
	comment := "Test Le Comment"

	scripts := "INSERT INTO comments(email, comment) VALUES (?, ?)"
	result, err := db.ExecContext(ctx, scripts, email, comment)
	if err != nil {
		panic(err)
	}

	insertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	fmt.Println("Sccess Insert New Comment with ID ", insertId)
}

func TestPrepareStatement(t *testing.T) {
	db := GetConnection()

	defer db.Close()

	ctx := context.Background()

	scripts := "INSERT INTO comments(email, comment) VALUES (?, ?)"
	statement, err := db.PrepareContext(ctx, scripts)
	if err != nil {
		panic(err)
	}

	defer statement.Close()

	for i := 0; i < 10; i++ {
		email := "rucco" + strconv.Itoa(i) + "@gmail"
		comment := "Ini komen ke " + strconv.Itoa(i)

		result, err := statement.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}

		insertId, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("Sccess Insert New Comment with ID ", insertId)
	}
}

func TestTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	scripts := "INSERT INTO comments(email, comment) VALUES (?, ?)"
	statement, err := db.PrepareContext(ctx, scripts)
	if err != nil {
		panic(err)
	}

	defer statement.Close()

	for i := 0; i < 10; i++ {
		email := "rucco" + strconv.Itoa(i) + "@gmail"
		comment := "Ini komen ke  " + strconv.Itoa(i)

		result, err := statement.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}

		insertId, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("Success Insert New Comment with ID ", insertId)
	}

	err = tx.Commit()
	if err != nil {
		panic(err)
	}
}
