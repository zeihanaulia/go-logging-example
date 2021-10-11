package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// EncoderConfig custome encoder
var EncoderConfig = zapcore.EncoderConfig{
	TimeKey:        "eventTime",
	LevelKey:       "severity",
	NameKey:        "logger",
	CallerKey:      "caller",
	MessageKey:     "message",
	StacktraceKey:  "stacktrace",
	LineEnding:     zapcore.DefaultLineEnding,
	EncodeLevel:    EncodeLevel,
	EncodeTime:     zapcore.ISO8601TimeEncoder,
	EncodeDuration: zapcore.SecondsDurationEncoder,
	EncodeCaller:   zapcore.ShortCallerEncoder,
}

var logLevelSeverity = map[zapcore.Level]string{
	zapcore.DebugLevel:  "DEBUG",
	zapcore.InfoLevel:   "INFO",
	zapcore.WarnLevel:   "WARNING",
	zapcore.ErrorLevel:  "ERROR",
	zapcore.DPanicLevel: "CRITICAL",
	zapcore.PanicLevel:  "ALERT",
	zapcore.FatalLevel:  "EMERGENCY",
}

func EncodeLevel(lv zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(logLevelSeverity[lv])
}

func StdConfig() zap.Config {
	cfg := zap.NewProductionConfig()
	cfg.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
	cfg.EncoderConfig = EncoderConfig
	cfg.Encoding = "json"
	cfg.OutputPaths = []string{"stdout"}
	cfg.ErrorOutputPaths = []string{"stderr"}
	cfg.DisableCaller = true
	return cfg
}

// Option context

var (
	logKeyServiceContext = "serviceContext"
)

type ServiceContext struct {
	Service string `json:"service,omitempty"`
	Version string `json:"version,omitempty"`
	Team    string `json:"team,omitempty"`
	Build   string `json:"build,omitempty"`
	Commit  string `json:"commit,omitempty"`
	Runtime string `json:"runtime,omitempty"`
}

func (s *ServiceContext) Clone() *ServiceContext {
	return &ServiceContext{
		Service: s.Service,
		Version: s.Version,
		Team:    s.Team,
		Build:   s.Build,
		Commit:  s.Commit,
		Runtime: s.Runtime,
	}
}

func (s *ServiceContext) MarshalLogObject(e zapcore.ObjectEncoder) error {
	e.AddString("service", s.Service)
	e.AddString("version", s.Version)
	e.AddString("team", s.Team)
	e.AddString("build", s.Build)
	e.AddString("commit", s.Commit)
	e.AddString("runtime", s.Runtime)
	return nil
}

// LogServiceContext add service name and version
func LogServiceContext(ctx *ServiceContext) zapcore.Field {
	return zap.Object(logKeyServiceContext, ctx)
}

// StdOptions should passing from env variable
func StdOptions() zap.Option {
	return zap.Fields(
		LogServiceContext(&ServiceContext{
			Service: "todo-service",
			Version: "v1.0.0",
			Team:    "platform",
			Build:   "54a162a.1",
			Commit:  "54a162a",
			Runtime: "go1.17.2",
		}),
	)
}
