package main

import (
	"net-up-test/internal"
)

func main() {
	api := internal.NewGinUsersAPI()
	api.RegisterUrls()
	api.Run()
}