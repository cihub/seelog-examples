package testlibrary

import (
	seelog "github.com/cihub/seelog"
	common "github.com/cihub/seelog/common"
	"io"
	"errors"
)

var logger seelog.LoggerInterface

func init() {
	DisableLog()
}

// DisableLog disables all library log output
func DisableLog() {
	logger = seelog.Disabled
}

// UseLogger uses a specified seelog.LoggerInterface to output library log.
// Use this func if you are using Seelog logging system in your app.
func UseLogger(newLogger seelog.LoggerInterface) {
	logger = newLogger
}

// SetLogWriter uses a specified io.Writer to output library log.
// Use this func if you are not using Seelog logging system in your app.
func SetLogWriter(writer io.Writer) error {
	if writer == nil {
		return errors.New("Nil writer")
	}
	
	newLogger, err := seelog.LoggerFromWriterWithMinLevel(writer, common.TraceLvl)
	if err != nil {
		return err
	}
	
	UseLogger(newLogger)
	return nil
}

// Call this before app shutdown
func FlushLog() {
	logger.Flush()
}