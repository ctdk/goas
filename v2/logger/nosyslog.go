// Tideland Go Application Support - Logger
//
// Copyright (C) 2012-2014 Frank Mueller / Tideland / Oldenburg / Germany
//
// All rights reserved. Use of this source code is governed
// by the new BSD license.

// +build windows plan9

package logger

import (
	"log"
)

// Dummy SysLogger for platforms without syslog
type SysLogger struct {
	
}

// A dummy NewSysLogger for platforms that don't support syslog. Only returns
// an error.
func NewSysLogger(tag string) (Logger, error) {
	log.Fatalf("Syslog not supported on this platform!")
	return &SysLogger{}, nil
}

// FAKE Debug logs a message at debug level.
func (sl *SysLogger) Debug(info, msg string) {
}

// FAKE Info logs a message at info level.
func (sl *SysLogger) Info(info, msg string) {
}

// FAKE Warning logs a message at warning level.
func (sl *SysLogger) Warning(info, msg string) {
}

// FAKE Error logs a message at error level.
func (sl *SysLogger) Error(info, msg string) {
}

// FAKE Critical logs a message at critical level.
func (sl *SysLogger) Critical(info, msg string) {
}

