package storage

import (
	"context"

	"github.com/jackc/pgx/v5"

	"simpleapp/internal/domain"
)

const userFields = ` name, surname, patronymic, email `

const (
	findUserQuery   = `SELECT` + userFields + `FROM users WHERE email = $1;`
	insertUserQuery = `INSERT INTO users (` + userFields + `) VALUES ($1, $2, $3, $4);`
)

type UserStorage interface {
	InsertUser(ctx context.Context, user domain.User) error
	FindUser(ctx context.Context, email string) (domain.User, error)
}

type userStorage struct {
	conn *pgx.Conn
}

func NewUserStorage(conn *pgx.Conn) UserStorage {
	return &userStorage{
		conn: conn,
	}
}

func (s *userStorage) FindUser(ctx context.Context, email string) (domain.User, error) {
	row := s.conn.QueryRow(ctx, findUserQuery, email)
	var name, surname, patronymic string
	err := row.Scan(
		&name,
		&surname,
		&patronymic,
		&email,
	)
	if err != nil {
		return domain.User{}, err
	}
	return domain.User{
		Name:       name,
		Surname:    surname,
		Patronymic: patronymic,
		Email:      email,
	}, nil
}

func (s *userStorage) InsertUser(ctx context.Context, user domain.User) error {
	_, err := s.conn.Exec(
		ctx,
		insertUserQuery,
		user.Name,
		user.Surname,
		user.Patronymic,
		user.Email,
	)
	return err
}
