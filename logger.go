// Copyright (c) 2020 E-Tiger Studio. All rights reserved.

package subr

import (
	"fmt"
	"log"
	"time"
)

const (
	//SubrPrefix = "[SUBR]"
	InfoPrefix = "INFO"
	WarningPrefix = "WARN"
	ErrorPrefix = "ERRR"
)

var (
	greenControlText   = string([]byte{27, 91, 57, 55, 59, 52, 50, 109})
	whiteControlText   = string([]byte{27, 91, 57, 48, 59, 52, 55, 109})
	yellowControlText  = string([]byte{27, 91, 57, 48, 59, 52, 51, 109})
	redControlText     = string([]byte{27, 91, 57, 55, 59, 52, 49, 109})
	blueControlText    = string([]byte{27, 91, 57, 55, 59, 52, 52, 109})
	magentaControlText = string([]byte{27, 91, 57, 55, 59, 52, 53, 109})
	cyanControlText    = string([]byte{27, 91, 57, 55, 59, 52, 54, 109})
	resetControlText   = string([]byte{27, 91, 48, 109})
)

type Logger struct {
	InstanceID
}

func getCommonPrefix() string {
	//return SubrPrefix + " " + time.Now().Format("2006/01/02 - 15:04:05")
	return " " + time.Now().Format("2006/01/02 - 15:04:05")
}

func getStatusPrefix(statusPrefix string, textColor string) string {
	//return " " + statusPrefix + " "
	return GetColoredText(" " + statusPrefix + " ", textColor)
}

func getInstancePrefix(id InstanceID) string {
	//return " " + string(id) + " "
	return GetColoredText(" " + string(id) + " ", blueControlText)
}

func GetColoredText(text string, textColor string) string {
	return textColor + text + resetControlText
}

func getLoggerForInstance(id InstanceID) *Logger {
	return &Logger{
		InstanceID: id,
	}
}

func (l *Logger) Infoln(v ...interface{}) {
	logln(getStatusPrefix(InfoPrefix, greenControlText), l.InstanceID, v...)
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.Infoln(fmt.Sprintf(format, v...))
	//logf(getStatusPrefix(InfoPrefix, greenControlText), l.InstanceID, format, v...)
}

func (l *Logger) Warnln(v ...interface{}) {
	logln(getStatusPrefix(WarningPrefix, yellowControlText), l.InstanceID, v...)
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	l.Warnln(fmt.Sprintf(format, v...))
	//logf(getStatusPrefix(WarningPrefix, yellowControlText), l.InstanceID, format, v...)
}

func (l *Logger) Errorln(v ...interface{}) {
	logln(getStatusPrefix(ErrorPrefix, redControlText), l.InstanceID, v...)
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.Errorln(fmt.Sprintf(format, v...))
	//logf(getStatusPrefix(ErrorPrefix, redControlText), l.InstanceID, format, v...)
}

func logln(statusPrefix string, id InstanceID, v ...interface{}) {
	log.Print(getCommonPrefix() + " |" + statusPrefix + "| |" + getInstancePrefix(id) + "| " + fmt.Sprintln(v...))
}

func logf(statusPrefix string, id InstanceID, format string, v ...interface{}) {
	logln(statusPrefix, id, fmt.Sprintf(format, v...))
	//fmt.Print(getCommonPrefix() + " |" + statusPrefix + "| |" + string(id) + "| " + fmt.Sprintf(format, v...))
}