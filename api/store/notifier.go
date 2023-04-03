package store

import (
	"log"
	"time"

	"github.com/lib/pq"

	"gitlab.com/khakibee/khakibee/api/logger"
)

const (
	BkngChangesPSQLChan = "bkng_changes_chan"
)

// listener aggregates the state of the listener connection.
type listener struct {
	*pq.Listener
	failed chan error
}

// NewListener creates a new listener for given PostgreSQL credentials.
func NewListener(connString string) *listener {
	l := &listener{failed: make(chan error, 2)}

	listener := pq.NewListener(
		connString,
		10*time.Second, time.Minute,
		l.LogListener)

	l.Listener = listener
	return l
}

// ListenAndNotify start listenig the given channel and runs callback function on notify
func (l *listener) ListenAndNotify(channelName string, callback func()) {
	if err := l.Listen(channelName); err != nil {
		l.Close()
		log.Fatal(err)
	}

	if err := l.listen(callback); err != nil {
		logger.EchoLogger.WithField("channel name", channelName).Error(err)
	}
}

// LogListener is the state change callback for the listener.
func (l *listener) LogListener(event pq.ListenerEventType, err error) {
	if err != nil {
		logger.EchoLogger.WithField("listener event", event).Error(err)
	}
	if event == pq.ListenerEventConnectionAttemptFailed {
		l.failed <- err
	}
}

// listen is the main loop of the listener to listen for notifications from the database.
func (l *listener) listen(callback func()) error {
	for {
		select {
		case <-l.Notify:
			callback()
		case err := <-l.failed:
			return err
		case <-time.After(time.Minute):
			go l.Ping()
		}
	}
}
