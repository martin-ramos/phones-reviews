package logs

import (
	"fmt"

	"go.uber.org/zap"
)

var sugar *zap.SugaredLogger

func InitLogger() error {
	// l, err := zap.NewDevelopment()
	l, err := zap.NewProduction()
	if err != nil {
		_ = fmt.Errorf("Can not create Zap logger %s", err.Error())
		return err
	}

	sugar = l.Sugar()
	return nil
}

func Log() *zap.SugaredLogger {
	return sugar
}
