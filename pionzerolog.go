package pionzerolog

import (
	"fmt"
	"github.com/ckb20110916/logusezerolog"
	"github.com/pion/logging"
	"github.com/rs/zerolog"
	"sync"
)

type logger struct {
	wrapper zerolog.Logger
}

func (l *logger) Trace(msg string) {
	l.wrapper.Trace().Msg(msg)
}

func (l *logger) Tracef(format string, args ...any) {
	l.wrapper.Trace().Msg(fmt.Sprintf(format, args...))
}

func (l *logger) Debug(msg string) {
	l.wrapper.Debug().Msg(msg)
}

func (l *logger) Debugf(format string, args ...any) {
	l.wrapper.Debug().Msg(fmt.Sprintf(format, args...))
}

func (l *logger) Info(msg string) {
	l.wrapper.Info().Msg(msg)
}

func (l *logger) Infof(format string, args ...any) {
	l.wrapper.Info().Msg(fmt.Sprintf(format, args...))
}

func (l *logger) Warn(msg string) {
	l.wrapper.Warn().Msg(msg)
}

func (l *logger) Warnf(format string, args ...any) {
	l.wrapper.Warn().Msg(fmt.Sprintf(format, args...))
}

func (l *logger) Error(msg string) {
	l.wrapper.Error().Msg(msg)
}

func (l *logger) Errorf(format string, args ...any) {
	l.wrapper.Error().Msg(fmt.Sprintf(format, args...))
}

var (
	Factory = &factory{}
)

type factory struct {
	mutex   sync.Mutex
	loggers []*logger
}

func (f *factory) NewLogger(scope string) logging.LeveledLogger {
	f.mutex.Lock()
	defer f.mutex.Unlock()
	l := &logger{
		wrapper: logusezerolog.Logger.With().Str("scope", scope).Logger(),
	}
	f.loggers = append(f.loggers, l)
	return l
}
