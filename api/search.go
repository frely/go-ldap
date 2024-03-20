package api

import (
	"log"

	"github.com/frely/go-ldap/api"
	"github.com/go-ldap/ldap/v3"
)

var res string

func Search(server string, bindUsername string, bindPassword string, baseDn string, searchRequestFilter string) string {
	l := api.Client(server, bindUsername, bindPassword)

	// Search for the given username
	searchRequest := ldap.NewSearchRequest(
		baseDn, // The base dn to search
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		searchRequestFilter,  // The filter to apply
		[]string{"dn", "cn"}, // A list attributes to retrieve
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
		//fmt.Printf("%s: %v\n", entry.DN, entry.GetAttributeValue("cn"))
		res = entry.DN + ": " + entry.GetAttributeValue("cn")
	}

	return res
}
