package helpers

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

const (
	LogsDirPath = "logs"
)

type IFileLogger interface {
	Info(message string)
	Warning(message string)
	Error(message string)
	Fatal(message string)
	GetLogger(prefix string) *log.Logger
}

type fileLogger struct {
	LogDirectory string
}

func (l *fileLogger) Info(message string) {
	log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.LUTC).Println(message)
	l.GetLogger("INFO: ").Println(message)
}

func (l *fileLogger) Warning(message string) {
	log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.LUTC).Println(message)
	l.GetLogger("WARNING: ").Println(message)
}

func (l *fileLogger) Error(message string) {
	log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.LUTC).Println(message)
	l.GetLogger("ERROR: ").Panicln(message)
}

func (l *fileLogger) Fatal(message string) {
	log.New(os.Stdout, "FATAL: ", log.Ldate|log.Ltime|log.LUTC).Println(message)
	l.GetLogger("FATAL: ").Fatalln(message)
}

func (l *fileLogger) GetLogger(prefix string) *log.Logger {
	getFilePath := l.setLogFile()
	return log.New(getFilePath, prefix, log.Ldate|log.Ltime|log.LUTC)
}

func (l *fileLogger) setLogFile() *os.File {
	year, month, day := time.Now().UTC().Date()
	monthDesc := strconv.Itoa(int(month))
	if len(monthDesc) < 2 {
		monthDesc += "0"
	}
	fileName := fmt.Sprintf("%v-%v-%v.log", year, monthDesc, day)
	filePath, _ := os.OpenFile(LogsDirPath+"/"+fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	return filePath
}

func NewFileLogger() *fileLogger {
	err := os.Mkdir(LogsDirPath, 0750)
	if err != nil {
		return nil
	}
	return &fileLogger{
		LogDirectory: LogsDirPath,
	}
}
