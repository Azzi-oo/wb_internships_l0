package main

import (
	"fmt"
	"time"
)

func sleep(milliseconds int) {
	duration := time.Duration(milliseconds) * time.Millisecond
	time.Sleep(duration)
}

func main() {
	fmt.Println("Начало программы")
	sleep(1000)
	fmt.Println("Прошло ... времени")
}
