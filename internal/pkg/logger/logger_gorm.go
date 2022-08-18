// @Author Ben.Zheng
// @DateTime 2022/8/18 9:52

package logger

// References:
// https://github.com/moul/zapgorm2/blob/master/zapgorm2.go

import (
	"context"
	"errors"
	"strings"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/benz9527/idk-doc/internal/pkg/consts"
	"github.com/benz9527/idk-doc/internal/pkg/intf"
)

type GormLogger struct {
	logger   *zap.SugaredLogger
	logLevel logger.LogLevel
}

func NewGormLogger(l *zap.SugaredLogger, cfgReader intf.IConfigurationReader) logger.Interface {
	env, err := cfgReader.GetString("app.env")
	if err != nil {
		env = consts.APP_RUNTIME_ENV_DEV
	}
	lvl, err := cfgReader.GetString("log.level")
	if err != nil || len(lvl) == 0 {
		lvl = consts.LogLevelCorrectOrDefault(lvl)
	}

	if strings.ToLower(env) == consts.APP_RUNTIME_ENV_PROD {
		return &GormLogger{
			logger:   l,
			logLevel: logger.Silent,
		}
	}
	// Default as dev output.
	return &GormLogger{
		logger: l,
		logLevel: func(lvl string) logger.LogLevel {
			switch strings.ToUpper(lvl) {
			case consts.APP_LOG_LVL_WARN:
				return logger.Warn
			case consts.APP_LOG_LVL_ERR:
				return logger.Error
			case consts.APP_LOG_LVL_INFO:
				return logger.Info
			default:
				return logger.Silent
			}
		}(lvl),
	}
}

func (g GormLogger) LogMode(level logger.LogLevel) logger.Interface {
	g.logLevel = level
	return g
}

func (g GormLogger) Info(ctx context.Context, format string, args ...any) {
	if g.logLevel < logger.Info {
		return
	}
	g.logger.Infof(format, args...)
}

func (g GormLogger) Warn(ctx context.Context, format string, args ...any) {
	if g.logLevel < logger.Warn {
		return
	}
	g.logger.Warnf(format, args...)
}

func (g GormLogger) Error(ctx context.Context, format string, args ...any) {
	if g.logLevel < logger.Error {
		return
	}
	g.logger.Errorf(format, args...)
}

func (g GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	if g.logLevel < logger.Silent {
		return
	}
	elapsed := time.Since(begin)
	sql, rows := fc()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		g.logger.Error("<TRACE ERROR>", zap.Error(err),
			zap.Duration("<ELAPSED>", elapsed),
			zap.Int64("<ROWS>", rows),
			zap.String("<SQL>", sql),
		)
		return
	}

	if elapsed > (100 * time.Millisecond) {
		g.logger.Warn("<TRACE WARN>", zap.Duration("elapsed", elapsed),
			zap.Int64("<ROWS>", rows),
			zap.String("<SQL>", sql),
		)
		return
	}

	g.logger.Debug("<TRACE>", zap.Duration("elapsed", elapsed),
		zap.Int64("<ROWS>", rows),
		zap.String("<SQL>", sql),
	)
}
