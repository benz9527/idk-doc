// @Author Ben.Zheng
// @DateTime 2022/8/11 14:03

package logger

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Test_logger_with_caller(t *testing.T) {
	asserter := assert.New(t)
	logger, err := zap.Config{
		Encoding: "json",
		Level:    zap.NewAtomicLevelAt(zapcore.DebugLevel),
		OutputPaths: []string{
			"stdout",
		},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:   "msg",
			LevelKey:     "lvl",
			EncodeLevel:  zapcore.CapitalLevelEncoder,
			TimeKey:      "ts",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			CallerKey:    "file",
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}.Build(zap.AddCaller())
	asserter.NotNil(logger)
	asserter.Nil(err)

	logger.Sugar().Debugf("This is a debug format, %s", "hello")
}
