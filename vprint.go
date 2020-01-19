package vprint

import (
	"fmt"
	"os"
	"strconv"
)

const (
	cBlack   = "\033[30m"
	cRed     = "\033[31m"
	cGreen   = "\033[32m"
	cYellow  = "\033[33m"
	cBlue    = "\033[34m"
	cMagenta = "\033[35m"
	cCyan    = "\033[36m"
	cWhite   = "\033[37m"
)

const (
	infoColor    = cBlue
	noticeColor  = cCyan
	warningColor = cYellow
	errorColor   = cRed
	debugColor   = cYellow
	endColor     = "\033[0m"
	endLine      = "\033[0m\n"
)

type vprintSettings struct {
	verbosity int64
}

func setup() vprintSettings {
	_settings := vprintSettings{
		verbosity: 0,
	}
	_settings.verbosity, _ = strconv.ParseInt(os.Getenv("VPRINTV"), 10, 16)
	return _settings
}

var settings = setup()

func Print(a ...interface{}) {
	if settings.verbosity > 0 {
		fmt.Fprint(os.Stderr, warningColor)
		fmt.Fprint(os.Stderr, a...)
		fmt.Fprint(os.Stderr, endColor)
	}
}
func Printf(format string, a ...interface{}) {
	if settings.verbosity > 0 {
		fmt.Fprint(os.Stderr, warningColor)
		fmt.Fprintf(os.Stderr, format, a...)
		fmt.Fprint(os.Stderr, endColor)
	}
}

func Red(a ...interface{}) string {
	return fmt.Sprint(cRed) + fmt.Sprint(a) + fmt.Sprint(endColor)
}
func Green(a ...interface{}) string {
	return fmt.Sprint(cGreen) + fmt.Sprint(a) + fmt.Sprint(endColor)
}
func Blue(a ...interface{}) string {
	return fmt.Sprint(cBlue) + fmt.Sprint(a) + fmt.Sprint(endColor)
}
