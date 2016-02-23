package httpblink
import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/saurabhsjoshi/gorpi/blink"
	"strconv"
	"encoding/json"
)

type status struct {
	ResponseCode int `json:"response_code"`
	Message string   `json:"message"`
}

type StatusMessage struct {
	Status status `json:"status"`
}

func newSuccessStatus() StatusMessage{
	return StatusMessage{
		Status: status{
			ResponseCode: 200,
			Message: "Query was successful!",
		}}
}

func newUnknownErrorStatus()  StatusMessage{
	return StatusMessage{
		Status: status{
			ResponseCode: 999,
			Message: "Unknown error!",
		}}
}

type Route struct {
	Name string
	Method string
	Pattern string
	HandlerFunc http.HandlerFunc
}

type Routes[] Route


func newRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = logger(handler, route.Name)

		router.
		Methods(route.Method).
		Path(route.Pattern).
		Name(route.Name).
		Handler(handler)
	}

	return router
}

var routes = Routes{
	Route{
		"toggleled",
		"GET",
		"/toggleled/",
		toggleLed,
	},
}

func toggleLed(w http.ResponseWriter, r *http.Request){
	pin,err := strconv.Atoi(r.FormValue("pin"))
	state,err2 := strconv.Atoi(r.FormValue("state"))


	if(err != nil || err2 != nil){
		json.NewEncoder(w).Encode(newUnknownErrorStatus())
		return
	}

	if(state == 0 || state == 1){
		blink.BlinkRelay(pin, state)
		json.NewEncoder(w).Encode(newSuccessStatus())
	}
}
