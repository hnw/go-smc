// +build !darwin

package smc

const KEY_CPU_TEMP string = ""

func Open() {
}

func GetTemperature(key string) float64 {
	return 0.0
}

func Close() {
}
