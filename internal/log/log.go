package log

import (
	"log"
	"os"
)

var (
    logger *log.Logger
)

func InitLogger() {
    logger = log.New(os.Stdout, "[verkurzen-api]", log.Default().Flags())
}

func Logger() *log.Logger {
    return logger
}
