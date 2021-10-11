# GO Logging Example

## Uber ZAP

```go
func main() {
    // call zap using company standard config
	logger, err := StdConfig().Build(StdOptions())
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	defer logger.Sync()

	sugar := logger.Sugar()
	svc := NewService(sugar)
	_ = svc.Todo(context.Background())
}
```

## Zerolog

```go
func main() {
	// call zap using company standard config
	log := zerolog.New(os.Stderr).
		With().
		Timestamp().
		Dict("serviceContext", ServiceContext()).
		Logger().Hook(SeverityHook{})

	svc := NewService(log)
	_ = svc.Todo(context.Background())
}
```