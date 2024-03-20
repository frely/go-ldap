package main

import (
	"fmt"
	"log"
	"os"

	"github.com/frely/go-ldap/api"
)

func main() {
	server, ok := os.LookupEnv("server")
	if !ok {
		log.Fatalln("server: variable is not defined")
	}

	bindUsername, ok := os.LookupEnv("bindUsername")
	if !ok {
		log.Fatalln("bindUsername: variable is not defined")
	}

	bindPassword, ok := os.LookupEnv("bindPassword")
	if !ok {
		log.Fatalln("bindPassword: variable is not defined")
	}

	baseDn, ok := os.LookupEnv("baseDn")
	if !ok {
		log.Fatalln("baseDn: variable is not defined")
	}

	searchRequestFilter, ok := os.LookupEnv("searchRequestFilter")
	if !ok {
		log.Fatalln("searchRequestFilter: variable is not defined")
	}

	// 搜索指定信息
	fmt.Println(api.Search(server, bindUsername, bindPassword, baseDn, searchRequestFilter))

}
