package user

import (
	"acme/model"
	"database/sql"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PostgresUserRepository struct {
	DB *sqlx.DB
}

func NewPostgresUserRepository(db *sqlx.DB) *PostgresUserRepository {
	return &PostgresUserRepository{DB: db}
}

func (repo *PostgresUserRepository) GetUsers() ([]model.User, error) {

	users := []model.User{}

	err := sqlx.Select(repo.DB, &users, "SELECT * FROM users")

	if err != nil {
		fmt.Println("Error querying the database:", err)
		return []model.User{}, errors.New("database could not be queried")
	}

	return users, nil
}

func (repo *PostgresUserRepository) AddUser(user model.User) (id int, err error) {

	err = repo.DB.QueryRow("INSERT INTO users (name) VALUES ($1) RETURNING id", user.Name).Scan(&id)

	if err != nil {
		fmt.Println("Error inserting user into the database:", err)
		return 0, errors.New("could not insert user")
	}

	return id, nil
}

func (repo *PostgresUserRepository) GetUser(id int) (user model.User, err error) {

	err = repo.DB.QueryRow("SELECT * FROM users WHERE id = $1", id).Scan(&user.ID, &user.Name)

	if err != nil {
		fmt.Println("Error retrieving user from the database:", err)
		return user, errors.New("could not retrieve user")
	}

	return user, nil
}

func (repo *PostgresUserRepository) DeleteUser(id int) error {

	result, err := repo.DB.Exec("DELETE FROM users WHERE id = $1", id)

	if err != nil {
		fmt.Println("Error deleting user from the database:", err)
		return errors.New("could not delete user")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return errors.New("user not found")
	}

	return nil
}

func (repo *PostgresUserRepository) UpdateUserName(id int, user *model.User) (updatedUser model.User, err error) {

	query := "UPDATE users SET name = $1 WHERE id = $2 RETURNING *"

	err = repo.DB.QueryRow(query, user.Name, id).Scan(&updatedUser.ID, &updatedUser.Name)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("user not found")
			return model.User{}, nil
		}
		fmt.Println("Error updating user name in the database:", err)
		return model.User{}, errors.New("could not update user name")
	}

	return updatedUser, nil
}

func (repo *PostgresUserRepository) Close() {

}
