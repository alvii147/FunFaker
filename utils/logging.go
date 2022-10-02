package utils

import (
	"log"
	"net/http"
	"os"
)

// loggers for each log level
var ErrorLogger *log.Logger = log.New(os.Stderr, "[E] ", log.Ldate|log.Ltime|log.Lmsgprefix)
var WarnLogger *log.Logger = log.New(os.Stdout, "[W] ", log.Ldate|log.Ltime|log.Lmsgprefix)
var InfoLogger *log.Logger = log.New(os.Stdout, "[I] ", log.Ldate|log.Ltime|log.Lmsgprefix)
var DebugLogger *log.Logger = log.New(os.Stdout, "[D] ", log.Ldate|log.Ltime|log.Lmsgprefix)

// log values with log level error
func LogError(v ...interface{}) {
	ErrorLogger.Println(v...)
}

// log values with log level warnings
func LogWarn(v ...interface{}) {
	WarnLogger.Println(v...)
}

// log values with log level information
func LogInfo(v ...interface{}) {
	InfoLogger.Println(v...)
}

// log values with log level debug
func LogDebug(v ...interface{}) {
	DebugLogger.Println(v...)
}

// log HTTP requests and responses data
func LogHTTPTraffic(request *http.Request, statusCode int) {
	if statusCode < 400 {
		LogInfo(request.Method, request.URL, request.Proto, statusCode, http.StatusText(statusCode))
	} else if statusCode < 500 {
		LogWarn(request.Method, request.URL, request.Proto, statusCode, http.StatusText(statusCode))
	} else {
		LogError(request.Method, request.URL, request.Proto, statusCode, http.StatusText(statusCode))
	}
}
