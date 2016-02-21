package httpblink

import (
	"log"
	"net/http"
)

func StartServer(port string){
	router := newRouter()
	log.Fatal(http.ListenAndServe(":" + port, router))
}
