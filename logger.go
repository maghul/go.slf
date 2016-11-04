package slf

import (
	"fmt"
)

/*
Logger contains two LoggerLevel outputs for Info
and Debug logging.
*/
type Logger struct {
	name        string
	description string

	lvl     Level
	own_lvl bool

	out        Output
	own_output bool

	parent      *Logger
	Info, Debug *LoggerLevel
}

func makeLogger(name string) *Logger {
	description := fmt.Sprint("Logger for ", name)
	l := &Logger{name: name, description: description, out: output_nilAdapter}
	l.initLogger()
	return l
}

func (l *Logger) initLogger() {
	l.Info = &LoggerLevel{l, Info}
	l.Debug = &LoggerLevel{l, Debug}
}

// Calculate the level based on parent levels.
func (l *Logger) calcLevel() Level {
	for l != nil {
		if l.own_lvl {
			return l.lvl
		}
		l = l.parent
	}
	return Off
}

/*
Get the name of the Logger
*/
func (l *Logger) Name() string {
	return l.name
}

/*
Get logger name, level and output of the Logger
*/
func (l *Logger) String() string {
	return fmt.Sprint("Logger:", l.Name(), ", level=", l.Level(), ", output=", fmt.Sprint(l.out))
}

/*
Get a brief description of the Logger
*/
func (l *Logger) Description() string {
	return l.description
}

/*
Change the description of the logger.
*/
func (l *Logger) SetDescription(d string) {
	l.description = d
}

/*
LoggerLevel is used to output logging at a specific level.
*/
type LoggerLevel struct {
	l   *Logger
	lvl Level
}

func (l *Logger) print(ref string, lvl Level, d ...interface{}) {
	l.out.Print(ref, lvl, d...)
}

func (l *Logger) printf(ref string, lvl Level, fmt string, d ...interface{}) {
	l.out.Printf(ref, lvl, fmt, d...)
}

/*
Output a log entry in the same way as fmt.Println
*/
func (lo *LoggerLevel) Println(d ...interface{}) {
	if lo.CanLog() {
		lo.l.print(lo.l.name, lo.lvl, d...)
	}
}

/*
Output a log entry in the same way as fmt.Printf
*/
func (lo *LoggerLevel) Printf(fmt string, d ...interface{}) {
	if lo.CanLog() {
		lo.l.printf(lo.l.name, lo.lvl, fmt, d...)
	}
}

/*
Check if the LoggerLevel will output any log. This us useful
if a log entry requires computation which is superfluous otherwise.
*/
func (lo *LoggerLevel) CanLog() bool {
	return lo.lvl <= lo.l.lvl
}
