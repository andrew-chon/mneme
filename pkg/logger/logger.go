package logger

import (
	"context"
	"fmt"
	"log/slog"
	"os"
)

// ContextLogger is intended to be used when we want logs to include request metadata
// when calling slog.Info/Warn/...Context(ctx), it will automatically extract specific
// context values and add it to the log attributes.  This contextual logging approach has some
// advantages compared to modifying a global instance of the logger with request attributes.  This allows for
// different requests being handled in different goroutines to have their own request metadata rather
// than sharing a global logging instance between goroutines that would clobber each others attributes.
type ContextLogger struct {
	*slog.JSONHandler
}

func (cl *ContextLogger) Handle(ctx context.Context, r slog.Record) error {
	newRecord := slog.NewRecord(r.Time, r.Level, r.Message, r.PC)

	// Add custom time attribute with 't' key and JavaScript ISO string format
	newRecord.AddAttrs(slog.String("t", r.Time.UTC().Format("2006-01-02T15:04:05.000Z")))

	// if correlationID, ok := ctx.Value(genesys.ContextCorrelationID).(string); ok {
	// 	newRecord.AddAttrs(slog.String(string(genesys.ContextCorrelationID), correlationID))
	// }

	// Add all other original attributes except the default time
	r.Attrs(func(a slog.Attr) bool {
		if a.Key != slog.TimeKey {
			newRecord.AddAttrs(a)
		}
		return true
	})

	return cl.JSONHandler.Handle(ctx, newRecord)
}

func NewContextLogger() *slog.Logger {
	var attrs []slog.Attr

	logLevel := slog.LevelInfo
	if level := os.Getenv("LOG_LEVEL"); level != "" {
		switch level {
		case "DEBUG":
			logLevel = slog.LevelDebug
		case "INFO":
			logLevel = slog.LevelInfo
		case "WARN":
			logLevel = slog.LevelWarn
		case "ERROR":
			logLevel = slog.LevelError
		default:
			fmt.Printf("LOG_LEVEL env value: [%s] is invalid, defaulting to INFO\n", level)
			logLevel = slog.LevelInfo
		}
	}

	handlerOpts := &slog.HandlerOptions{
		Level: logLevel,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			// Remove the default time attribute - we'll add our custom one
			if a.Key == slog.TimeKey {
				return slog.Attr{}
			}
			return a
		},
	}

	baseHandler := slog.NewJSONHandler(os.Stdout, handlerOpts)

	if len(attrs) > 0 {
		baseHandler = baseHandler.WithAttrs(attrs).(*slog.JSONHandler)
	}

	contextHandler := &ContextLogger{JSONHandler: baseHandler}

	return slog.New(contextHandler)
}

func SetGlobalContextLogger() {
	logger := NewContextLogger()
	slog.SetDefault(logger)
}
