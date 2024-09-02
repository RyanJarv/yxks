package utils

import (
	"context"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

const (
	colorScheme = "pastel28"

	Red   Color = "\033[31m"
	Green Color = "\033[32m"
	Cyan  Color = "\033[36m"
	Gray  Color = "\033[37m"

	ErrorLogLevel LogLevel = iota
	InfoLogLevel
	DebugLogLevel
)

type Color string

func (c Color) Color(s ...string) string {
	return string(c) + strings.Join(s, " ") + "\033[0m"
}

type LogLevel int

func NewContext(parentCtx context.Context) Context {
	ctx := Context{
		Context: parentCtx,
		Error:   log.New(os.Stderr, Red.Color("[ERROR] "), 0),
		Info:    log.New(os.Stdout, Green.Color("[INFO] "), 0),
		Debug:   log.New(os.Stdout, Gray.Color("[DEBUG] "), 0),
	}

	ctx.Debug.SetOutput(io.Discard)
	return ctx
}

type Context struct {
	context.Context
	LogLevel LogLevel
	Error    *log.Logger
	Info     *log.Logger
	Debug    *log.Logger
}

func (ctx *Context) SetLoggingLevel(level LogLevel) Context {
	ctx.LogLevel = level

	if int(level) >= int(ErrorLogLevel) {
		ctx.Error = log.New(os.Stderr, Red.Color("[ERROR] "), 0)
	} else {
		ctx.Error.SetOutput(io.Discard)
	}

	if int(level) >= int(InfoLogLevel) {
		ctx.Info = log.New(os.Stderr, Green.Color("[INFO] "), 0)
	} else {
		ctx.Info.SetOutput(io.Discard)
	}

	if int(level) >= int(DebugLogLevel) {
		ctx.Debug = log.New(os.Stderr, Gray.Color("[DEBUG] "), 0)
	} else {
		ctx.Info.SetOutput(io.Discard)
	}
	return *ctx
}

func (ctx Context) WithCancel() (Context, context.CancelFunc) {
	var cancel context.CancelFunc
	ctx.Context, cancel = context.WithCancel(ctx.Context)
	return Context{
		Context: ctx,
		Info:    ctx.Info,
		Debug:   ctx.Debug,
		Error:   ctx.Error,
	}, cancel
}

func (ctx Context) IsRunning(msg ...string) bool {
	select {
	case <-ctx.Done():
		if len(msg) != 0 {
			ctx.Info.Println(msg)
		}
		return false
	default:
		return true
	}
}

func (ctx Context) IsDone(msg ...string) bool {
	return !ctx.IsRunning(msg...)
}

func (ctx Context) Sleep(delay time.Duration) {
	select {
	case <-ctx.Done():
	case <-time.After(delay):
	}
}
