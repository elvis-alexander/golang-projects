package main

import (
	"strings"
)

/*
	-> make chat room
	-> join chat room
	-> say something
	-> leave chat room
	-> brodcast to all chat rooms
	-> view all my chat rooms
	-> view all users
	-> number of active users
*/
type Gender string

const (
	Male   Gender = "male"
	Female Gender = "gender"
)

type user struct {
	/*unique identifier*/
	username string
	gender   Gender
	age      int
}

type room struct {
	peopleMap map[string]*string
	admin     string
}

type chat struct {
	/*room name to room*/
	roomMap map[string]*room
	/*all users logged on*/
	users   []user
}

func (this *chat) numUsers() int {
	return len(this.users)
}

func (this *chat) allUsers() string {
	return strings.Join(this.users, ",")
}

/*returns true if able to create new room*/
func (this *chat) newRoom(roomName string, adminName string) bool {
	_, contains := this.roomMap[roomName]
	if contains {
		return false
	}
	this.roomMap[adminName] = &room{admin:adminName}
	return true
}

func (this *chat) joinRoom()

func main() {

}
