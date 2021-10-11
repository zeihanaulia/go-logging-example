package main

import (
	"context"
	"errors"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"

	"github.com/zeihanaulia/go-logging-example/xerrors"
)

func main() {
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	log := zerolog.New(os.Stderr).
		With().
		Timestamp().
		Dict("serviceContext", ServiceContext()).
		Logger().Hook(SeverityHook{})

	svc := NewService(log)
	_ = svc.Todo(context.Background())
}

type service struct {
	log zerolog.Logger
}

func NewService(log zerolog.Logger) service {
	return service{log}
}

func (s service) Todo(ctx context.Context) error {
	if err := s.process(ctx); err != nil {
		s.log.Error().Stack().Msg(err.Error())
	}
	return nil
}

func (s service) process(ctx context.Context) error {
	return xerrors.WrapErrorf(errors.New("something wrong"), xerrors.ErrorCodeUnknown, "process")
}
