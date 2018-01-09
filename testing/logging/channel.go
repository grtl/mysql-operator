package logging

import (
	"github.com/sirupsen/logrus"
)

// Hook is a testing hook, which publishes all entries to the channel.
type Hook struct {
	entries chan *logrus.Entry
}

const entriesChannelBufferSize = 100

// NewGlobal installs a test hook for the global logger.
func NewGlobal() *Hook {
	hook := new(Hook)
	hook.entries = make(chan *logrus.Entry, entriesChannelBufferSize)
	logrus.AddHook(hook)

	return hook
}

// Fire is run whenever an Entry is logged.
func (t *Hook) Fire(e *logrus.Entry) error {
	t.entries <- e
	return nil
}

// Levels returns levels on which hook should be called.
func (t *Hook) Levels() []logrus.Level {
	return logrus.AllLevels
}

// GetEntriesChan returns entries channel.
func (t *Hook) GetEntriesChan() <-chan *logrus.Entry {
	return t.entries
}
