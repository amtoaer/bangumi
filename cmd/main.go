package main

import (
	"bangumi/internal/login"
	"fmt"
)

func main() {
	api := login.Login("bgm16655f7fd3916a663", "47b62b4cf43cc36716753e83eb13725a", "5212").NewSession()
	result, err := api.UserInfo("amtoaer")
	if err != nil {
		return
	}
	fmt.Println(result)
}
