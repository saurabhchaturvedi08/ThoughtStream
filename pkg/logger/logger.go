package logger

import (
	"log"
)

func Info(msg string) {
	log.Printf("ℹ️  %s\n", msg)
}

func Error(err error, msg string) {
	log.Printf("❌ %s: %v\n", msg, err)
}
