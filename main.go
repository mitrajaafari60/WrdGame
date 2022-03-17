//main of project
package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"strconv"
	"wordDemo/controller"
)

var port int

//setArgs can change the listen port by input flags such as -port=80
func setArgs(args []string) {
	flag.IntVar(&port, "port", 8000, "limit the number of parallel requests with flag 'port'")
	flag.Parse()
}

func main() {
	setArgs(os.Args)
	log.Println("Lisen Port ", port)

	// in case of no arguments http server is lunched
	//details of port and path wasn't defined, so defaults are used
	http.HandleFunc("/start", controller.GetWord)
	http.HandleFunc("/check", controller.CalculateInput)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), nil))
}
