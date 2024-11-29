package utils

import (
	"errors"
	"os"

	"golang.org/x/crypto/ssh"
)

var SSH_Config *ssh.ClientConfig

func InitializeConnection() error {
	var hostKey ssh.PublicKey
	key, err := os.ReadFile(SSH_Key)
	if err != nil {
		return errors.New("cannot read ssh key")
	}

	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		return errors.New("unable to parse key")
	}

	SSH_Config = &ssh.ClientConfig{
		User: SSH_USER,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.FixedHostKey(hostKey),
	}
	return nil
}