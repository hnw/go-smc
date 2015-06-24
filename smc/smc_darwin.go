package smc

/*
#cgo LDFLAGS: -framework IOKit
#include "_smc.c"
*/
import "C"

/*
Temperature sensors
-------------------
TB0T
TC0D
TC0P
TM0P
TN0P
Th0H
Ts0P
TN1P
Th1H
*/
const KEY_CPU_TEMP string = "TC0P"

func Open() {
	C.smc_init()
}

func GetTemperature(key string) float64 {
	return float64(C.SMCGetTemperature(C.CString(key)))
}

func Close() {
	C.smc_close()
}
