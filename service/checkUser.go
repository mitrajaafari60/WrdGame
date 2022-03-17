package service

import (
	"wordDemo/DataModel"
)

// CheckUserValidation check if the word is assigned to this user
func CheckUserValidation(user string) bool {

	data := DataModel.UsersData[user]
	//	log.Println("data:", data)
	if len(data) != DataModel.WordsLen {
		return false
	}
	return true
}
