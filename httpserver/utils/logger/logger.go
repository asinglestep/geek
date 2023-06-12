package logger

import (
	"context"
	"os"
	"strings"

	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	gLogger = NewLogger()
)

type Logger struct {
	*zap.SugaredLogger
}

func SetLogger(log *Logger) {
	gLogger = log
}

func NewLogger(opts ...Option) *Logger {
	opt := newOptions(opts...)

	cfg := zapcore.EncoderConfig{
		CallerKey:    "caller",
		EncodeCaller: zapcore.ShortCallerEncoder,

		TimeKey:    "time",
		EncodeTime: MillisTimeEncoder,

		LevelKey:    "level",
		EncodeLevel: UppercaseLevelEncoder,

		MessageKey:     "msg",
		NameKey:        "logger",
		EncodeDuration: zapcore.StringDurationEncoder,
	}

	level := zap.NewAtomicLevelAt(ToZapLevel(opt.Level))

	log := &Logger{
		SugaredLogger: zap.New(
			zapcore.NewCore(
				zapcore.NewConsoleEncoder(cfg),
				zapcore.NewMultiWriteSyncer(
					zapcore.Lock(zapcore.AddSync(os.Stdout)),
				),
				level,
			),
		).WithOptions(zap.AddCaller(), zap.AddCallerSkip(1)).Sugar(),
	}

	return log
}

func ToZapLevel(level string) zapcore.Level {
	switch strings.ToLower(level) {
	case "info":
		return zapcore.InfoLevel

	case "warn":
		return zapcore.WarnLevel

	case "error":
		return zapcore.ErrorLevel

	default:
		return zapcore.DebugLevel
	}
}

func Debug(args ...interface{}) {
	gLogger.Debug(args...)
}

func Info(args ...interface{}) {
	gLogger.Info(args...)
}

func Warn(args ...interface{}) {
	gLogger.Warn(args...)
}

func Error(args ...interface{}) {
	gLogger.Error(args...)
}

func Debugf(format string, args ...interface{}) {
	gLogger.Debugf(format, args...)
}

func Infof(format string, args ...interface{}) {
	gLogger.Infof(format, args...)
}

func Warnf(format string, args ...interface{}) {
	gLogger.Warnf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	gLogger.Errorf(format, args...)
}

func With(args ...interface{}) *Logger {
	return &Logger{SugaredLogger: gLogger.SugaredLogger.With(args...).Desugar().WithOptions(zap.AddCallerSkip(-1)).Sugar()}
}

func WithContext(ctx context.Context) *Logger {
	return With("trace-id", trace.SpanContextFromContext(ctx).TraceID())
}
