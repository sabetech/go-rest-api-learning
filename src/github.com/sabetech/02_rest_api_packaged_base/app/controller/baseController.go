package baseController

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var router *mux.Router

func initHandlers() {

}

func Start() {
	router = mux.NewRouter()

	initHandlers()
	fmt.Printf("router initialized and listening on 8002\n")
	log.Fatal(http.ListenAndServe(":8002", router))
}
