package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type User struct {
	ID    int
	Name  string
	Email string
}

func main() {
	ctx := context.Background()

	// CONNECTION
	db, err := pgxpool.New(
		ctx,
		"postgres://postgres:password@localhost:5432/mydb",
	)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// CREATE TABLE
	_, err = db.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS users(
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL,
			email TEXT UNIQUE NOT NULL
		)
	`)
	if err != nil {
		log.Fatal(err)
	}

	// CREATE
	userID, err := createUser(
		ctx,
		db,
		"John Doe",
		"john@example.com",
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Created User ID:", userID)

	// READ ONE
	user, err := getUserByID(
		ctx,
		db,
		userID,
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("User: %+v\n", user)

	// UPDATE
	err = updateUser(
		ctx,
		db,
		userID,
		"John Updated",
		"updated@example.com",
	)
	if err != nil {
		log.Fatal(err)
	}

	updatedUser, _ := getUserByID(
		ctx,
		db,
		userID,
	)

	fmt.Printf("Updated User: %+v\n", updatedUser)

	// READ ALL
	users, err := getAllUsers(
		ctx,
		db,
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("All Users:")
	for _, u := range users {
		fmt.Printf("%+v\n", u)
	}

	// DELETE
	err = deleteUser(
		ctx,
		db,
		userID,
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("User deleted")
}

// CREATE
func createUser(
	ctx context.Context,
	db *pgxpool.Pool,
	name string,
	email string,
) (int, error) {

	var id int

	err := db.QueryRow(
		ctx,
		`
		INSERT INTO users(name,email)
		VALUES($1,$2)
		RETURNING id
		`,
		name,
		email,
	).Scan(&id)

	return id, err
}

// READ ONE
func getUserByID(
	ctx context.Context,
	db *pgxpool.Pool,
	id int,
) (User, error) {

	var user User

	err := db.QueryRow(
		ctx,
		`
		SELECT id,name,email
		FROM users
		WHERE id=$1
		`,
		id,
	).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
	)

	return user, err
}

// READ ALL
func getAllUsers(
	ctx context.Context,
	db *pgxpool.Pool,
) ([]User, error) {

	rows, err := db.Query(
		ctx,
		`
		SELECT id,name,email
		FROM users
		ORDER BY id
		`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User

		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
		)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, rows.Err()
}

// UPDATE
func updateUser(
	ctx context.Context,
	db *pgxpool.Pool,
	id int,
	name string,
	email string,
) error {

	_, err := db.Exec(
		ctx,
		`
		UPDATE users
		SET
			name=$1,
			email=$2
		WHERE id=$3
		`,
		name,
		email,
		id,
	)

	return err
}

// DELETE
func deleteUser(
	ctx context.Context,
	db *pgxpool.Pool,
	id int,
) error {

	_, err := db.Exec(
		ctx,
		`
		DELETE FROM users
		WHERE id=$1
		`,
		id,
	)

	return err
}
