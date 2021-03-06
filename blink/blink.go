package blink
/*
#cgo LDFLAGS: -lwiringPi
#include <wiringPi.h>

void setupPin(int pin){
	wiringPiSetup();
	pinMode(pin,OUTPUT);
}

void toggleLED(int state){
	if(state == 1){
		digitalWrite(0, HIGH);
	} else {
		digitalWrite(0,LOW);
	}
}
*/
import "C"

func Blink(pin int, state int){
	//Only using 0 for now
	C.setupPin(C.int(0))
	C.toggleLED(C.int(state))
}

func BlinkRelay(pin int, state int){
	//1 is on and 0 is off for the relay
	if(state == 0){
		Blink(pin, 1)
	} else {
		Blink(pin, 0)
	}

}