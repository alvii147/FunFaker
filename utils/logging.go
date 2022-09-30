package utils

import (
	"log"
	"net/http"
)

// log levels
type LogLevel string

const (
	LogLevelError LogLevel = "ERROR"
	LogLevelWarn  LogLevel = "WARN"
	LogLevelInfo  LogLevel = "INFO"
	LogLevelDebug LogLevel = "DEBUG"
)

// log values with given log level
func LogWithLevel(level LogLevel, v ...interface{}) {
	logPrefix := "[" + string(level[0]) + "]"
	logSlice := append([]interface{}{logPrefix}, v...)
	log.Println(logSlice...)
}

// log values with log level error
func LogError(v ...interface{}) {
	LogWithLevel(LogLevelError, v...)
}

// log values with log level warnings
func LogWarn(v ...interface{}) {
	LogWithLevel(LogLevelWarn, v...)
}

// log values with log level information
func LogInfo(v ...interface{}) {
	LogWithLevel(LogLevelInfo, v...)
}

// log values with log level debug
func LogDebug(v ...interface{}) {
	LogWithLevel(LogLevelDebug, v...)
}

// log HTTP requests and responses data
func LogHTTPTraffic(request *http.Request, statusCode int, err error) {
	if statusCode < 400 {
		LogInfo(request.Method, request.URL, request.Proto, statusCode, http.StatusText(statusCode))
	} else if statusCode < 500 {
		LogWarn(request.Method, request.URL, request.Proto, statusCode, http.StatusText(statusCode), err)
	} else {
		LogError(request.Method, request.URL, request.Proto, statusCode, http.StatusText(statusCode), err)
	}
}
