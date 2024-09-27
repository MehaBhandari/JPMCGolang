package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type User struct {
	ID   int
	Name string
}

type Associates struct {
	Id        string
	Name      string
	CreatedAt string
	Avatar    string
}

type UserService interface {
	GetUser(id int) string
}

func (userService Service) InitUsers() {
	response, _ := http.Get("http://localhost:7000/associates")
	byteData, _ := io.ReadAll(response.Body)
	json.Unmarshal(byteData, userService.associateArray)
}

func (userService Service) GetUser(userId int) string {
	return userService.associateArray[0].Name
}

// Test This Function
func GetUserDetails(service UserService, id int) string {
	abc := service.GetUser(id)
	abc = abc + "dfysgfiu"
	return abc
}

func (userService GetUserDetails) GetUserDetails(userId int) string {
	return userService.associateArray[0].Name
}

type Service struct {
	associateArray []Associates
}

func main() {
}
