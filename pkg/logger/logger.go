package logger

import (
	"log"
	"os"
)

func New() *log.Logger {
	return log.New(
		os.Stdout,
		"[myapp]",
		log.LstdFlags|log.Lshortfile,
	)
}
