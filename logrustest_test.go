package logrustest

import (
    "errors"
    "testing"
    "time"
)

func TestPasses(t *testing.T) {
    logger := New(t)
    logger.WithField("x", "y").Warn("everything went well, you shouldn't see this except with go test -v")
}

func TestFails(t *testing.T) {
    logger := New(t)
    logger.Debug("you should see this")
    logger.WithError(errors.New("bad error")).Error("this test fails so you should see this")
    t.Fail()
}

func BenchmarkPasses(b *testing.B) {
    logger := New(b)
    logger.WithField("n", b.N).Info("you should see this with go test -bench=. -run=BenchmarkPasses")
    for n := 0; n < b.N; n++ {
        time.Sleep(time.Millisecond)
    }
}
