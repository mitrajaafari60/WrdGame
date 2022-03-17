//controller for checking word
package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"wordDemo/DataModel"
	"wordDemo/service"
)

// CalculateInput returns array of colores base on definition of the game Gary,Green,Yellow for each alphabet
func CalculateInput(w http.ResponseWriter, r *http.Request) {

	agent := r.Header.Get("user-Agent")
	if r.Method == http.MethodOptions {
		return
	}
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	Body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	defer r.Body.Close()

	//	log.Println(Body, agent)
	var vData DataModel.GameInput
	err = json.Unmarshal(Body, &vData)
	log.Println("method /check called with data:", string(Body), err)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//	log.Println("len:", len(agent), len(vData.Word), DataModel.WordsLen)
	if len(agent) <= 0 || len(vData.Word) != DataModel.WordsLen {
		log.Println("Error in length of input:", len(vData.Word))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !service.CheckUserValidation(agent) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	response := service.CheckWord(vData.Word, agent)

	log.Println("response:", response)
	json.NewEncoder(w).Encode(response)
}
