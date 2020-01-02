package logrustest

import (
	"strings"
	"testing"

	"github.com/sirupsen/logrus"
)

// New creates a logrus debug logger which writes to testing.T.Log.
func New(t testing.TB) *logrus.Logger {
	return &logrus.Logger{
		Out:          Writer{t},
		Level:        logrus.DebugLevel,
		ReportCaller: true,
		Formatter: &logrus.TextFormatter{
			DisableTimestamp:       true,
			DisableLevelTruncation: true,
		},
	}
}

// Writer adapts the testing.T/testing.B interface to the io.Writer interface.
type Writer struct {
	testing.TB
}

func (t Writer) Write(data []byte) (int, error) {
	t.Log(strings.TrimSuffix(string(data), "\n"))
	return len(data), nil
}
