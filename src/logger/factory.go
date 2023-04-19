package logger

import (
	"go.uber.org/zap"
)

type Factory interface {
	Channel(channel string) (*zap.Logger, error)
	RegisterLogger(channel string, loggerResolver func() (*zap.Logger, error))
}
