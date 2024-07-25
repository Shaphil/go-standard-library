package main

import (
	"log"
	"os"
)

var (
	WarniLogger *log.Logger
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
	FatalLogger *log.Logger
)

func init() {
	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	InfoLogger = log.New(file, "INFO: ", log.LUTC|log.Lmicroseconds|log.Llongfile)
	WarniLogger = log.New(file, "WARNING: ", log.LUTC|log.Lmicroseconds|log.Llongfile)
	ErrorLogger = log.New(file, "ERROR: ", log.LUTC|log.Lmicroseconds|log.Llongfile)
	FatalLogger = log.New(file, "FATAL: ", log.LUTC|log.Lmicroseconds|log.Llongfile)
}

func main() {
	InfoLogger.Println("This is info")
	WarniLogger.Println("Oh no, a warning!")
	ErrorLogger.Println("Error! Please Stop")
	FatalLogger.Println("Program Quit")
}
