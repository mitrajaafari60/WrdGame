package service

import (
	"log"
	"math/rand"
	"time"
	"wordDemo/DataModel"
)

//SetRandomWord assign one random  word from dictionary to the user-agent and save it localy
func SetRandomWord() string {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	randomIndex := r1.Intn(len(DataModel.Datas))
	log.Println("randomIndex:", randomIndex, len(DataModel.Datas))
	return DataModel.Datas[randomIndex]
}
