package logger

import "go.uber.org/zap"

// New returns a development logger when local==true, production logger otherwise.
// Mirrors the legacy InitializeLogger behaviour.
func New(local bool) *zap.Logger {
	if local {
		l, err := zap.NewDevelopment()
		if err != nil {
			panic(err)
		}
		return l
	}
	l, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	return l
}
