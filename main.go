package main

import (
	"os"
	"github.com/hxline/pgmuxsem-master/service"
)

// const (
// 	TEST_DB_HOST     = "10.17.50.48"
// 	TEST_DB_PORT     = "5432"
// 	TEST_DB_USERNAME = "postgres"
// 	TEST_DB_PASSWORD = "postgres"
// 	TEST_DB_NAME     = "hxline"
// )
func main() {
	a := service.App{}
	// a.Initialize(TEST_DB_USERNAME, TEST_DB_PASSWORD, TEST_DB_NAME)

	a.Initialize(
		os.Getenv("APPa_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))

	a.Run(":8080")
}
