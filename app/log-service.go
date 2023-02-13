package main

import (
	"fmt"
	"log"
	"os"
)

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
	FatalLogger   *log.Logger
)

func infoLogger(message string) {
	file, err := os.OpenFile(getLogFile(), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(logFileError(err))
	}

	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	InfoLogger.Println(message)
}

func warningLogger(message string) {
	file, err := os.OpenFile(getLogFile(), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(logFileError(err))
	}

	WarningLogger = log.New(file, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger.Println(message)
}

func errorLogger(message string) {
	file, err := os.OpenFile(getLogFile(), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(logFileError(err))
	}

	ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger.Println(message)
}

func fatalLogger(message string) {
	file, err := os.OpenFile(getLogFile(), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(logFileError(err))
	}

	FatalLogger = log.New(file, "FATAL: ", log.Ldate|log.Ltime|log.Lshortfile)
	FatalLogger.Println(message)
}

func logFileError(err error) string {
	return fmt.Sprintf("Could not access calculator.log file: err: %v", err)
}

func getLogFile() string {
	
	return "./log/calculator.log"
}