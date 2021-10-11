package main

import (
	"context"
	"errors"
	"log"

	"github.com/zeihanaulia/go-logging-example/xerrors"
	"go.uber.org/zap"
)

func main() {
	logger, err := StdConfig().Build(StdOptions())
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	defer logger.Sync()

	sugar := logger.Sugar()
	svc := NewService(sugar)
	_ = svc.Todo(context.Background())
}

type service struct {
	log *zap.SugaredLogger
}

func NewService(log *zap.SugaredLogger) service {
	return service{log}
}

func (s service) Todo(ctx context.Context) error {
	if err := s.process(ctx); err != nil {
		s.log.Error(zap.Error(err))
	}
	return nil
}

func (s service) process(ctx context.Context) error {
	return xerrors.WrapErrorf(errors.New("something wrong"), xerrors.ErrorCodeUnknown, "process")
}
