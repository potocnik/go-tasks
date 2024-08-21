package logging

import (
	"log"
	"os"
)

var (
	debugLogger *log.Logger
	infoLogger  *log.Logger
	errorLogger *log.Logger
)

func SetUp() {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("Could not open file logs.txt for logging")
	}
	log.SetOutput(file)
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Llongfile)
	debugLogger = log.New(file, "[DEBUG]: ", log.LstdFlags|log.Lmicroseconds|log.Llongfile)
	infoLogger = log.New(file, "[INFO]:", log.LstdFlags)
	errorLogger = log.New(file, "[ERROR]:", log.LstdFlags|log.Lmicroseconds|log.Llongfile)
}

func Info(message string) {
	infoLogger.Println(message)
}

func Debug(message string, data any) {
	debugLogger.Println(message, data)
}

func Error(message string, err error) {
	errorLogger.Println(message, err)
}
