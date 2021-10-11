package main

import (
	"github.com/rs/zerolog"
)

//ServiceContext this service context should be provide by env var
func ServiceContext() *zerolog.Event {
	svcCtx := zerolog.Dict()
	svcCtx.Str("service", "todo-service")
	svcCtx.Str("version", "v1.0.0")
	svcCtx.Str("team", "platform")
	svcCtx.Str("build", "54a162a.1")
	svcCtx.Str("commit", "54a162a")
	svcCtx.Str("runtime", "go1.17.2")
	return svcCtx
}

type SeverityHook struct{}

func (h SeverityHook) Run(e *zerolog.Event, level zerolog.Level, msg string) {
	if level != zerolog.NoLevel {
		e.Str("severity", logLevelSeverity[level])
	}
}

var logLevelSeverity = map[zerolog.Level]string{
	zerolog.DebugLevel: "DEBUG",
	zerolog.InfoLevel:  "INFO",
	zerolog.WarnLevel:  "WARNING",
	zerolog.ErrorLevel: "ERROR",
	zerolog.PanicLevel: "ALERT",
	zerolog.FatalLevel: "EMERGENCY",
}
