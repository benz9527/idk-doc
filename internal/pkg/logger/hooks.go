// @Author Ben.Zheng
// @DateTime 2022/8/11 10:33

package logger

import "go.uber.org/zap/zapcore"

type ZapEntryHook func(entry zapcore.Entry) error

// Through this hook, we can intercept log before it handled
// by zap.
func syslogInterceptor(entry zapcore.Entry) error {
	go func(e zapcore.Entry) {
		// TODO(Ben) For extension.
	}(entry)
	return nil
}
