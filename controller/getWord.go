//controller for starting Game an assigning word for userAgent in backend to be solved!
package controller

import (
	"log"
	"net/http"
	"wordDemo/DataModel"
	"wordDemo/service"
)

// GetWord API returns responseCode 200 in case of Ok, otherwise BadRequest
func GetWord(w http.ResponseWriter, r *http.Request) {

	agent := r.Header.Get("user-Agent")
	if r.Method == http.MethodOptions {
		return
	}

	if len(agent) <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	DataModel.UsersData[agent] = service.SetRandomWord()
	log.Print("for user ", agent)
	log.Println(" word:", DataModel.UsersData[agent])
	w.WriteHeader(http.StatusOK)
}
