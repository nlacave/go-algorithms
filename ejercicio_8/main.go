package main

import "fmt"

func main() {
	f := flipsMinimos("100110")
	fmt.Println(f)
}

func flipsMinimos(pwd string) int32 {
	if len(pwd)%2 == 0 {
		cont := 0
		for i := 0; i < len(pwd); i += 2 {
			if pwd[i] != pwd[i+1] {
				cont++
			}
		}
		return int32(cont)
	}
	return -1
}
