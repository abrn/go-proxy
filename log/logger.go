package log

import (
	"github.com/gookit/color"
)

var Logger *ColorLogger

var (
	inner = color.HEXStyle("#cca4ef", "#212121")
	outer = color.HEXStyle("#5d1ec9", "#212121")
	info = color.HEXStyle("#693c72")
	infoPre = color.HEXStyle("#693c72", "#1e1121")
	debug = color.HEXStyle("#9c5cdb")
	debugPre = color.HEXStyle("#9c5cdb", "#200836")
	trace = color.HEXStyle("#00c3d9")
	tracePre = color.HEXStyle("#00c3d9", "#002b30")
	warn = color.HEXStyle("#d97642")
	warnPre = color.HEXStyle("#d97642", "#33190b")
	fatal = color.HEXStyle("#ed002b")
	fatalPre = color.HEXStyle("#ed002b", "#240007")
)

// ColorLogger - logs Output in color with optional verbosity
type ColorLogger struct {
	VeryVerbose bool
	Verbose     bool
	Prefix      string
}

// Trace - write a blue trace message at max verbosity
func (l ColorLogger) Trace(f string, args ...interface{}) {
	if !l.VeryVerbose {
		return
	}
	color.Printf("%s%s ", l.getPrefix(), tracePre.Sprint(" TRACE >"))
	trace.Printf(f, args...)
}

// Debug - write a green debug message when verbose
func (l ColorLogger) Debug(f string, args ...interface{}) {
	if !l.Verbose {
		return
	}
	color.Printf("%s%s ", l.getPrefix(), debugPre.Sprint(" DEBUG >"))
	debug.Printf(f, args...)
}

// Info - write a green informational message
func (l ColorLogger) Info(f string, args ...interface{}) {
	color.Printf("%s%s ", l.getPrefix(), infoPre.Sprint(" INFO  >"))
	info.Printf(f, args...)
}

// Warn - write a yellow warning message
func (l ColorLogger) Warn(f string, args ...interface{}) {
	color.Printf("%s%s ", l.getPrefix(), warnPre.Sprint(" WARN  >"))
	warn.Printf(f, args...)
}

// Error - write a red error message
func (l ColorLogger) Error(f string, args ...interface{}) {
	color.Printf("%s%s ", l.getPrefix(), fatalPre.Sprint(" ERROR >"))
	fatal.Printf(f, args...)
}

func (l ColorLogger) getPrefix() string {
	if l.Prefix != "" {
		return l.Prefix
	}
	l.Prefix = color.Sprintf("%s%s%s", outer.Sprint("["), inner.Sprint("rn-proxy"), outer.Sprint("]"))
	return l.Prefix
}

// NoLogger - used for logging nothing while keeping code functional
type NoLogger struct{}

func (l NoLogger) Trace(f string, args ...interface{}) {}
func (l NoLogger) Debug(f string, args ...interface{}) {}
func (l NoLogger) Info(f string, args ...interface{})  {}
func (l NoLogger) Warn(f string, args ...interface{})  {}
func (l NoLogger) Error(f string, args ...interface{}) {}

// InterfaceLogger - the main interface for each logger type
type InterfaceLogger interface {
	Trace(f string, args ...interface{})
	Debug(f string, args ...interface{})
	Info(f string, args ...interface{})
	Warn(f string, args ...interface{})
	Error(f string, args ...interface{})
}