package models

import (
	"errors"

	"github.com/bogle2921/elenktis/api/internal/database"
	"github.com/bogle2921/elenktis/api/internal/utils"
)

type User struct {
	ID       int64
	Admin    bool
	Username string `binding:"required"`
	Password string `binding:"required"`
	Salt     []byte
}

func (u *User) AddUser() error {
	query := "insert into users (admin, username, password, salt) values (?, ?, ?, ?)"
	stmt, err := database.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	u.Salt = utils.GenerateSalt()
	u.Password = utils.HashPassword(u.Salt, u.Password)
	res, err := stmt.Exec(u.Admin, u.Username, u.Password, u.Salt)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	u.ID = id
	return err
}

func (u *User) Validate() error {
	query := "select id, password, salt from users where username = ?"
	row := database.DB.QueryRow(query, u.Username)

	var returnedPass string
	var returnedSalt []byte
	err := row.Scan(&u.ID, &returnedPass, &returnedSalt)
	if err != nil {
		return errors.New("invalid credentials")
	}

	validPassword := utils.VerifyHash(u.Password, returnedPass, returnedSalt)
	if !validPassword {
		// bad password
		return errors.New("invalid credentials")
	}

	return nil
}
