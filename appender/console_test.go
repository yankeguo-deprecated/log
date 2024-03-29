package appender

import (
	"bytes"
	"github.com/novakit/log/event"
	"github.com/novakit/log/labels"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestConsole(t *testing.T) {
	e := event.Event{
		Timestamp: time.Date(2011, 11, 11, 11, 11, 11, 0, time.UTC),
		Project:   "test",
		Env:       "test",
		Hostname:  "test",
		Topic:     "test",
		Labels:    labels.Labels{"test": "test"},
		Message:   "test",
	}
	b := &bytes.Buffer{}
	a := Console(b)
	_ = a.Log(e)
	require.Equal(t, "2011-11-11T11:11:11.000+0000 [test] {\"test\":\"test\"} test\n", b.String())
}
