package log

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"time"
)

type TextFormatter struct {
	TimeFormat string
}

func (f *TextFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer

	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	b.WriteString(fmt.Sprintf("[%s]: ", entry.Level.String()))

	if f.TimeFormat == "" {
		f.TimeFormat = time.RFC3339
	}
	b.WriteString(entry.Time.UTC().Format(f.TimeFormat))

	if entry.Message != "" {
		b.WriteString(fmt.Sprintf(": %s", entry.Message))
	}

	for key, val := range entry.Data {
		b.WriteString(fmt.Sprintf(" (%s: %v)", key, val))
	}
	b.WriteByte('\n')
	return b.Bytes(), nil
}
