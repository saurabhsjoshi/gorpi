package main

import (
	"github.com/saurabhsjoshi/gorpi/blink"
	"github.com/saurabhsjoshi/gorpi/httpblink"
	"os"
	"flag"
	"fmt"
)

func main(){
	f_blink := flag.Bool("blink", false, "Toggles the LED")
	f_http := flag.Bool("http", false, "Starts http server")
	f_port := flag.String("port", "8080", "HTTP port (Optional)")
	f_pin := flag.Int("pin", -1, "Pin number")
	f_pinState := flag.Int("state", -1, "Pin state 0 = OFF and 1 = ON")
	flag.Parse()

	if(*f_blink == true){
		if(*f_pin != -1 && *f_pinState != -1){
			if(*f_pinState != 0 || *f_pinState != 1){
				blink.Blink(*f_pin,*f_pinState)
			} else {
				fmt.Println("Incorrect value for -state")
				os.Exit(0)
			}


		} else {
			fmt.Println("Please provide -pin and -state")
			os.Exit(0)
		}

	} else if(*f_http == true){
		httpblink.StartServer(*f_port)
	}


}

