package service

import (
	"log"
	"strings"
	"wordDemo/DataModel"
)

//CheckWord check 5 Alphabet with assigned word to response based on defined colores Gray,Green,Yellow
func CheckWord(word, agent string) (resp DataModel.GameResponse) {
	defer func() {
		if re := recover(); re != nil {
			log.Println("Recovered in CheckWord", re)
			resp.Status = DataModel.ERROR
		}
	}()
	data := DataModel.UsersData[agent]

	if len(data) != DataModel.WordsLen {
		resp.Status = DataModel.ERROR
		return
	}

	log.Println("check want , userData:", data, word)
	okOnes := 0
	for i := 0; i < len(data); i++ {

		//	log.Println("checking:", data[i], word[i])
		if data[i] == word[i] {
			resp.Value[i] = DataModel.Green
			okOnes++
		} else if strings.Contains(word[i:i+1], data) {
			resp.Value[i] = DataModel.Yellow
		} else {
			resp.Value[i] = DataModel.Grey
		}
	}

	if okOnes == DataModel.WordsLen {
		resp.Status = DataModel.SUCCESS
	} else {
		resp.Status = DataModel.FAILED
	}
	return
}
