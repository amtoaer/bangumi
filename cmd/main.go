package main

import (
	"bangumi/internal/login"
	"encoding/json"
	"fmt"
)

func main() {
	api := login.Login("bgm16655f7fd3916a663", "47b62b4cf43cc36716753e83eb13725a", "5212").NewSession()
	result, _ := api.UserProgress("amtoaer", "")
	content, _ := json.Marshal(result)
	fmt.Println(string(content))
}
