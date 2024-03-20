package api

import (
	"fmt"
	"log"

	"github.com/go-ldap/ldap"
)

func GetUserInfo() {
	// Search for the given username
	searchRequest := ldap.NewSearchRequest(
		"DC=wlyd,DC=local", // The base dn to search
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		"(sAMAccountName=wangxiao)", // The filter to apply
		[]string{"dn", "cn"},        // A list attributes to retrieve
		nil,
	)

	sr, err := l.Search(searchRequest)
	if err != nil {
		log.Fatal(err)
	}
	if len(sr.Entries) != 1 {
		log.Fatal("User does not exist or too many entries returned")
	}

	for _, entry := range sr.Entries {
		fmt.Printf("%s: %v\n", entry.DN, entry.GetAttributeValue("cn"))
	}
}
