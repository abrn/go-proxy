package main

import (
	"fmt"
	"github.com/mgutz/ansi"
)

// Logger - the main interface for each logger type
type Logger interface {
	Trace(f string, args ...interface{})
	Debug(f string, args ...interface{})
	Info(f string, args ...interface{})
	Warn(f string, args ...interface{})
	Error(f string, args ...interface{})
}

// NullLogger - used for logging nothing while keeping code functional
type NullLogger struct {}

func (l NullLogger) Trace(f string, args ...interface{}) {}
func (l NullLogger) Debug(f string, args ...interface{}) {}
func (l NullLogger) Info(f string, args ...interface{}) {}
func (l NullLogger) Warn(f string, args ...interface{}) {}
func (l NullLogger) Error(f string, args ...interface{}) {}


// ColorLogger - logs output in color with optional verbosity
type ColorLogger struct {
	VeryVerbose		bool
	Verbose 		bool
	Prefix			string
	Color			bool
}

// Trace - write a blue trace message at max verbosity
func (l ColorLogger) Trace(f string, args ...interface{}) {
	if !l.VeryVerbose {
		return
	}
	l.output("blue", f, args...)
}

// Debug - write a green debug message when verbose
func (l ColorLogger) Debug(f string, args ...interface{}) {
	if !l.Verbose {
		return
	}
	l.output("green", f, args...)
}

// Info - write a green informational message
func (l ColorLogger) Info(f string, args ...interface{}) {
	l.output("green", f, args...)
}

// Warn - write a yellow warning message
func (l ColorLogger) Warn(f string, args ...interface{}) {
	l.output("yellow", f, args...)
}

// Error - write a red error message
func (l ColorLogger) Error(f string, args ...interface{}) {
	l.output("red", f, args...)
}

// output - the main function used for writing text
func (l ColorLogger) output(color, f string, args ...interface{}) {
	if l.Color && color != "" {
		f = ansi.Color(f, color)
	}
	fmt.Printf(fmt.Sprintf("%s%s\n", l.Prefix, f), args...)
}


type PacketLogger struct {

}
