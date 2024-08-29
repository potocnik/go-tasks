package logging

import (
	"log"
	"os"
)

func SetUpLogging() {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("Could not open file logs.txt for logging")
	}
	log.SetOutput(file)
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Llongfile)
}

func Info(message string) {
	log.Println("[INFO]: " + message)
}

func Debug(message string, data any) {
	log.Println("[DEBUG]: "+message, data)
}

func Error(message string, err error) {
	log.Println("[ERROR]: "+message, err)
}
