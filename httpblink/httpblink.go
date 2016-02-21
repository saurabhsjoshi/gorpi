package httpblink

import (
	"log"
	"net/http"
	"fmt"
)

func StartServer(port string){
	router := newRouter()
	fmt.Println("Starting server on port " + port)
	log.Fatal(http.ListenAndServe(":" + port, router))
}
