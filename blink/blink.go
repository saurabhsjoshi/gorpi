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
	C.setupPin(C.int(pin))
	C.toggleLED(C.int(state))
}