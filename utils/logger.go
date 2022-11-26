package utils

import (
	"log"
	"go.uber.org/zap"
)

// for dependency injection
func NewLogger() *zap.SugaredLogger {
	logger, err := zap.NewProduction()
	if(err != nil){
		log.Fatal(err)
	}
	return logger.Sugar()
}
