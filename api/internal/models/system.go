package models

import (
	"bytes"
	"errors"

	"github.com/bogle2921/elenktis/api/internal/database"
	"github.com/bogle2921/elenktis/api/internal/utils"
	"golang.org/x/crypto/ssh"
)

type System struct {
	ID           int64
	IP           string `binding:"required"`
	FriendlyName string
	Services     []string
}

func (s *System) AddSystem() error {
	query := "insert into systems (ip, friendlyname, services) values (?, ?, [?])"
	stmt, err := database.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(s.IP, s.FriendlyName, s.Services)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	s.ID = id
	return err
}

func (s *System) RunCommand(cmd string) (string, error) {
	client, err := ssh.Dial("tcp", s.IP+":22", utils.SSH_Config)
	if err != nil {
		return "", errors.New("cannot connect to " + s.IP)
	}

	session, err := client.NewSession()
	if err != nil {
		return "", errors.New("cannot open session to " + s.IP)
	}
	defer session.Close()

	var stdout bytes.Buffer
	session.Stdout = &stdout
	err = session.Run(cmd)

	return stdout.String(), err
}
