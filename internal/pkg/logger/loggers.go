// @Author Ben.Zheng
// @DateTime 2022/8/10 16:48

package logger

// https://pkg.go.dev/go.uber.org/zap#pkg-overview
// Self-defined zap logger/sugarLogger pls read upper URL.

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/benz9527/idk-doc/internal/pkg/intf"
)

const (
	LOGGER_JSON  = "JSON"
	LOGGER_PLAIN = "PlainText"
)

type LoggerFactory func(reader intf.IConfigurationReader) *zap.SugaredLogger

var (
	Loggers = map[string]LoggerFactory{
		LOGGER_JSON:  newLoggerBy(LOGGER_JSON),
		LOGGER_PLAIN: newLoggerBy(LOGGER_PLAIN),
	}
)

func NewLogger(cfgReader intf.IConfigurationReader) *zap.SugaredLogger {
	logStyle, err := cfgReader.GetString("log.style")
	if err != nil || len(logStyle) == 0 {
		// Set default logger as JSON style.
		logStyle = LOGGER_JSON
	}

	fn, ok := Loggers[logStyle]
	if !ok {
		// Set default logger as JSON logger.
		fn = newLoggerBy(LOGGER_JSON)
	}
	return fn(cfgReader)
}

func newLoggerBy(logStyle string) LoggerFactory {
	return func(cfgReader intf.IConfigurationReader) *zap.SugaredLogger {
		var (
			logger *zap.Logger
			cores  = make([]zapcore.Core, 0, 2)
			outCfg fileOutConfig
		)

		outCfg = newOutConfig(logStyle, cfgReader)
		if consoleCore, err := outCfg.consoleOutConfig.build(); err == nil {
			cores = append(cores, consoleCore)
		}
		if fileCores, err := outCfg.build(); err == nil {
			cores = append(cores, fileCores...)
		}

		if len(cores) == 0 {
			cfg := zap.NewDevelopmentEncoderConfig()
			zapcore.NewCore(
				zapcore.NewJSONEncoder(cfg),
				zapcore.Lock(os.Stdout),
				zapcore.DebugLevel,
			)
		}

		// Caller should be added here instead of getting them from core.
		// Multiple clients output.
		logger = zap.New(
			zapcore.NewTee(cores...),
			zap.AddCaller(),
			zap.AddStacktrace(zap.ErrorLevel),
		)
		return logger.Sugar()
	}
}
