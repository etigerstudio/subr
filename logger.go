// Copyright (c) 2020 E-Tiger Studio. All rights reserved.

package subr

import (
	"fmt"
)

const (
	InfoPrefix = "Info"
	WarningPrefix = "Warning"
	ErrorPrefix = "Error"
)

func Infoln(v ...interface{}) {
	logln(InfoPrefix, v...)
}

func Infof(format string, v ...interface{}) {
	logf(InfoPrefix, format, v...)
}

func Warnln(v ...interface{}) {
	logln(WarningPrefix, v...)
}

func Warnf(format string, v ...interface{}) {
	logf(WarningPrefix, format, v...)
}

func Errorln(v ...interface{}) {
	logln(ErrorPrefix, v...)
}

func Errorf(format string, v ...interface{}) {
	logf(ErrorPrefix, format, v...)
}

func logln(prefix string, v ...interface{}) {
	fmt.Print(prefix + " " + fmt.Sprintln(v...))
}

func logf(prefix string, format string, v ...interface{}) {
	fmt.Print(prefix + " " + fmt.Sprintf(format, v...))
}