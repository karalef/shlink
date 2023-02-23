package cmd

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

// Getenv returns the value of the environment variable named by the key.
// If the variable is empty, the defaultValue is returned.
func Getenv(name string, defaultValue ...string) string {
	value := os.Getenv(name)
	if value == "" && len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return value
}

// ListenSig listens for SIGTERM and SIGINT signals.
// If the signal is received, onCancel is called and then the function returns.
// If the context is cancelled, the function returns immediately.
func ListenSig(ctx context.Context, onCancel func(os.Signal)) {
	close := make(chan os.Signal, 1)
	signal.Notify(close, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(close)

	var s os.Signal
	select {
	case s = <-close:
		onCancel(s)
	case <-ctx.Done():
	}
}
