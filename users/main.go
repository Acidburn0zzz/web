package main

import "github.com/zqzca/web/users/api"

func main() {
	api := api.NewUserAPI("zqz-users-dev", "dylan", 5000)
	api.App.Listen()
}
