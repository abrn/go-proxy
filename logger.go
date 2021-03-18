package main

import (
	"fmt"
	"github.com/mgutz/ansi"
)

// LogInterface - the main interface for each logger type
type LogInterface interface {
	Trace(f string, args ...interface{})
	Debug(f string, args ...interface{})
	Info(f string, args ...interface{})
	Warn(f string, args ...interface{})
	Error(f string, args ...interface{})
	Output(color string, f string, args ...interface{})
}

// ColorLogger - logs Output in color with optional verbosity
type ColorLogger struct {
	VeryVerbose bool
	Verbose     bool
	Prefix      string
	Color       bool
}

// Trace - write a blue trace message at max verbosity
func (l ColorLogger) Trace(f string, args ...interface{}) {
	if !l.VeryVerbose {
		return
	}
	l.Output("blue", f, args...)
}

// Debug - write a green debug message when verbose
func (l ColorLogger) Debug(f string, args ...interface{}) {
	if !l.Verbose {
		return
	}
	l.Output("green", f, args...)
}

// Info - write a green informational message
func (l ColorLogger) Info(f string, args ...interface{}) {
	l.Output("green", f, args...)
}

// Warn - write a yellow warning message
func (l ColorLogger) Warn(f string, args ...interface{}) {
	l.Output("yellow", f, args...)
}

// Error - write a red error message
func (l ColorLogger) Error(f string, args ...interface{}) {
	l.Output("red", f, args...)
}

// Output - printf with an ANSI color
func (l ColorLogger) Output(color, f string, args ...interface{}) {
	if l.Color && color != "" {
		f = ansi.Color(f, color)
	}
	fmt.Printf(fmt.Sprintf("%s%s\n", l.Prefix, f), args...)
}

func (l ColorLogger) getPrefix() string {
	if l.Prefix != "" {
		return l.Prefix
	}
	return ""
}

// NoLogger - used for logging nothing while keeping code functional
type NoLogger struct{}

func (l NoLogger) Trace(f string, args ...interface{})                {}
func (l NoLogger) Debug(f string, args ...interface{})                {}
func (l NoLogger) Info(f string, args ...interface{})                 {}
func (l NoLogger) Warn(f string, args ...interface{})                 {}
func (l NoLogger) Error(f string, args ...interface{})                {}
func (l NoLogger) Output(color string, f string, args ...interface{}) {}
