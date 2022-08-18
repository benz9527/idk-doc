// @Author Ben.Zheng
// @DateTime 2022/8/11 9:21

package logger

import (
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/benz9527/idk-doc/internal/pkg/consts"
	"github.com/benz9527/idk-doc/internal/pkg/intf"
)

type consoleOutConfig struct {
	encoderKey string
	lvl        zapcore.Level
}

func (c consoleOutConfig) build() (zapcore.Core, error) {
	var core zapcore.Core
	cfg := zapcore.EncoderConfig{
		MessageKey:    "msg",
		LevelKey:      "lvl",
		EncodeLevel:   zapcore.CapitalLevelEncoder,
		TimeKey:       "ts",
		EncodeTime:    zapcore.ISO8601TimeEncoder,
		CallerKey:     "file",
		EncodeCaller:  zapcore.ShortCallerEncoder,
		StacktraceKey: "stack",
	}

	consoleLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= c.lvl
	})
	core = zapcore.NewCore(
		ternaryEncoder(c.encoderKey, cfg),
		zapcore.AddSync(os.Stdout),
		consoleLevel,
	)
	return core, nil
}

type fileOutConfig struct {
	consoleOutConfig
	enable     bool
	cout       string // @field common output path short name. Save non-error level log.
	eout       string // @field error output path short name. Save error level log. Default is "stderr".
	maxMbSize  int
	maxBackups int
	maxDays    int
	compress   bool
}

func (c fileOutConfig) build() ([]zapcore.Core, error) {
	if !c.enable {
		return nil, errors.New("disabled")
	}

	var (
		cores = make([]zapcore.Core, 0, 2)
	)

	cfg := zapcore.EncoderConfig{
		MessageKey:    "msg",
		LevelKey:      "lvl",
		EncodeLevel:   zapcore.CapitalColorLevelEncoder,
		TimeKey:       "ts",
		EncodeTime:    zapcore.ISO8601TimeEncoder,
		CallerKey:     "caller",
		EncodeCaller:  zapcore.ShortCallerEncoder,
		StacktraceKey: "stack",
	}

	coutLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.WarnLevel && lvl >= c.lvl
	})
	eoutLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.WarnLevel && lvl >= c.lvl
	})

	cores = append(cores, zapcore.NewCore(
		ternaryEncoder(c.encoderKey, cfg),
		zapcore.AddSync(c.rotateWriter(c.cout)),
		coutLevel,
	))
	cores = append(cores, zapcore.NewCore(
		ternaryEncoder(c.encoderKey, cfg),
		zapcore.AddSync(c.rotateWriter(c.eout)),
		eoutLevel,
	))

	return cores, nil
}

func (c fileOutConfig) rotateWriter(filepath string) io.Writer {
	return &lumberjack.Logger{
		Filename:   filepath,
		MaxSize:    c.maxMbSize,
		MaxAge:     c.maxDays,
		MaxBackups: c.maxBackups,
		LocalTime:  true,
		Compress:   c.compress,
	}
}

func newOutConfig(encoderStyle string, cfgReader intf.IConfigurationReader) fileOutConfig {
	var (
		cout, eout string
		lvl        zapcore.Level
	)

	enable, err := cfgReader.GetBoolean("log.file.enable")
	if err != nil {
		enable = false
	}

	env, err := cfgReader.GetString("app.env")
	if err != nil || len(env) == 0 ||
		env == consts.APP_RUNTIME_ENV_DEV {
		cout, eout = filepath.Join(os.TempDir(), "idk.log"), filepath.Join(os.TempDir(), "idk_err.log")
	} else {
		// Product env.
		rwd, _ := cfgReader.GetString(consts.APP_ROOT_WORKING_DIR)

		// Handle log output path missing.
		if path, err := cfgReader.GetString("log.file.common_out_path"); err != nil || len(path) == 0 {
			cout = filepath.Join(rwd, "log", "idk.log")
		} else {
			cout = path
		}

		if path, err := cfgReader.GetString("log.file.err_out_path"); err != nil || len(path) == 0 {
			cout = filepath.Join(rwd, "log", "idk_err.log")
		} else {
			cout = path
		}
	}

	level, _ := cfgReader.GetString("log.level")
	lvl = getLevel(level, env)

	mb, err := cfgReader.GetInt64("log.file.max_mb_size")
	if err != nil || mb == 0 {
		mb = int64(10)
	}

	days, err := cfgReader.GetInt64("log.file.max_days")
	if err != nil || days == 0 {
		days = int64(16)
	}

	backups, err := cfgReader.GetInt64("log.file.max_backups")
	if err != nil || backups == 0 {
		backups = int64(7)
	}

	compress, err := cfgReader.GetBoolean("log.file.compress")
	if err != nil {
		compress = false
	}

	return fileOutConfig{
		consoleOutConfig: consoleOutConfig{
			lvl:        lvl,
			encoderKey: encoderStyle,
		},
		enable:     enable,
		cout:       cout,
		eout:       eout,
		maxMbSize:  int(mb),
		maxDays:    int(days),
		maxBackups: int(backups),
		compress:   compress,
	}
}

func ternaryEncoder(condition string, cfg zapcore.EncoderConfig) zapcore.Encoder {
	if condition == consts.APP_LOG_ENC_PLAIN {
		return zapcore.NewConsoleEncoder(cfg)
	}
	return zapcore.NewJSONEncoder(cfg)
}

func getLevel(lvl, env string) zapcore.Level {
	switch strings.ToUpper(lvl) {
	case consts.APP_LOG_LVL_DEBUG:
		return zapcore.DebugLevel
	case consts.APP_LOG_LVL_INFO:
		return zapcore.InfoLevel
	case consts.APP_LOG_LVL_WARN:
		return zapcore.WarnLevel
	case consts.APP_LOG_LVL_ERR:
		return zapcore.ErrorLevel
	default:
		if env == consts.APP_RUNTIME_ENV_PROD {
			return zapcore.InfoLevel
		}
		return zapcore.DebugLevel
	}
}
