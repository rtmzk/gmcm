package log

import (
	"github.com/marmotedu/errors"
	"gmcm/pkg/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"strconv"
)

var (
	name          = utils.GetEnv("log.name", "gmcm")
	level         = utils.GetEnv("log.level", "info")
	format        = utils.GetEnv("log.format", "console")
	color, _      = strconv.ParseBool(utils.GetEnv("log.enable-color", "true"))
	discaller, _  = strconv.ParseBool(utils.GetEnv("log.disable-caller", "true"))
	outputPath    = []string{utils.GetEnv("log.output-path", "stdout")}
	erroutputPath = []string{utils.GetEnv("log.error-output-path", "stderr")}
)

type zapLogger struct {
	zaplogger *zap.Logger
}

func NewZapLogger() (*zapLogger, error) {
	var zapLevel zapcore.Level
	if err := zapLevel.UnmarshalText([]byte(level)); err != nil {
		zapLevel = zapcore.InfoLevel
	}

	encodeLevel := zapcore.CapitalLevelEncoder
	if format == "console" && color {
		encodeLevel = zapcore.CapitalColorLevelEncoder
	}

	encoderConfig := zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "timestamp",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    encodeLevel,
		EncodeTime:     timeEncoder,
		EncodeDuration: milliSecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	loggerConfig := &zap.Config{
		Level:             zap.NewAtomicLevelAt(zapLevel),
		Development:       false,
		DisableCaller:     discaller,
		DisableStacktrace: color,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding:         format,
		EncoderConfig:    encoderConfig,
		OutputPaths:      outputPath,
		ErrorOutputPaths: erroutputPath,
	}

	var err error
	l, err := loggerConfig.Build(zap.AddStacktrace(zapcore.PanicLevel), zap.AddCallerSkip(1))
	if err != nil {
		return nil, errors.New("zap logger build constructs failed.")
	}
	logger := &zapLogger{
		zaplogger: l.Named(name),
	}

	zap.RedirectStdLog(l)
	return logger, nil
}
