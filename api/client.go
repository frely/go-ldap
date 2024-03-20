package api

import (
	"crypto/tls"
	"log"

	"github.com/go-ldap/ldap/v3"
)

func Client(server string, bindUsername string, bindPassword string) {
	l, err := ldap.Dial("tcp", server)
	if err != nil {
		log.Fatal("Failed to connect to server:\n", err)
	}
	defer l.Close()

	// Reconnect with TLS
	err = l.StartTLS(&tls.Config{InsecureSkipVerify: true})
	if err != nil {
		log.Fatal(err)
	}

	// First bind with a read only user
	err = l.Bind(bindUsername, bindPassword)
	if err != nil {
		log.Fatal(err)
	}

}
