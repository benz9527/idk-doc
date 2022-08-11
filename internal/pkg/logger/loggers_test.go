// @Author Ben.Zheng
// @DateTime 2022/8/11 13:18

package logger

import (
	"errors"
	"testing"

	"github.com/benz9527/idk-doc/internal/pkg/file"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func Test_json_logger(t *testing.T) {
	asserter := assert.New(t)

	t.Parallel()
	t.Run("json without file out", func(tt *testing.T) {
		tt.Parallel()
		v := viper.New()
		v.Set("log.file.enable", false)
		reader := file.NewSimpleReader(v)
		logger := newLoggerBy(LOGGER_JSON)(reader)
		defer logger.Sync()
		asserter.NotNil(logger)
		// format out put.
		logger.Debugf("This is a debug format, %s", "hello without file out")
		logger.Errorf("This is an error format with stack trace, %v", errors.New("an accident without file out"))
	})

	t.Run("json with file out", func(tt *testing.T) {
		tt.Parallel()
		v := viper.New()
		v.Set("log.file.enable", true)
		reader := file.NewSimpleReader(v)
		logger := newLoggerBy(LOGGER_JSON)(reader)
		defer logger.Sync()
		asserter.NotNil(logger)
		// format out put.
		logger.Debugf("This is a debug format, %s", "hello with file out")
		logger.Errorf("This is an error format with stack trace, %v", errors.New("an accident with file out"))
	})
}

func Test_plaintext_logger(t *testing.T) {
	asserter := assert.New(t)

	t.Parallel()
	t.Run("json without file out", func(tt *testing.T) {
		tt.Parallel()
		v := viper.New()
		v.Set("log.file.enable", false)
		reader := file.NewSimpleReader(v)
		logger := newLoggerBy(LOGGER_PLAIN)(reader)
		defer logger.Sync()
		asserter.NotNil(logger)
		// format out put.
		logger.Debugf("This is a debug format, %s", "hello without file out")
		logger.Errorf("This is an error format with stack trace, %v", errors.New("an accident without file out"))
	})

	t.Run("json with file out", func(tt *testing.T) {
		tt.Parallel()
		v := viper.New()
		v.Set("log.file.enable", true)
		reader := file.NewSimpleReader(v)
		logger := newLoggerBy(LOGGER_PLAIN)(reader)
		defer logger.Sync()
		asserter.NotNil(logger)
		// format out put.
		logger.Debugf("This is a debug format, %s", "hello with file out")
		logger.Errorf("This is an error format with stack trace, %v", errors.New("an accident with file out"))
	})
}
