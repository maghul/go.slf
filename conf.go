package slf

import (
	"errors"
	"fmt"
)

/*
Set a parent logger. This is used to group and control output
and levels. If a logger has a parent set it will output through
the parent output logger by default. If the level of the logger is
set to Parent then logging levels will be controlled by the parent.
*/
func (l *Logger) SetParent(parent *Logger) {
	l.parent = parent
	scanSetLevels()
	scanSetOutputLoggers()
}

/*
Get the parent logger.
*/
func (l *Logger) Parent() *Logger {
	return l.parent
}

/*
Set the level of a logger by name. Returns
error if this failed. This is a convienence function.
*/
func SetLevel(loggerName, levelName string) error {
	logger, err := FindLogger(loggerName)
	if err != nil {
		return errors.New(fmt.Sprint("No logger named '", loggerName, "':", err))
	}
	level, err := FindLevel(levelName)
	if err != nil {
		return errors.New(fmt.Sprint("No log level named '", levelName, "':", err))
	}

	logger.SetLevel(level)
	return nil
}

/*
Get the Level which the Logger will log at. The levels returned will
be Off, Info, and Debug.
*/
func (l *Logger) Level() Level {
	return l.lvl
}

/*
Set the Level which the Logger will log at. The levels are
Off, Info, Debug and Parent.
*/
func (l *Logger) SetLevel(lvl Level) {
	if lvl == Parent {
		l.own_lvl = false
	} else {
		l.own_lvl = true
		l.lvl = lvl
	}

	scanSetLevels()
}

func scanSetLevels() {
	loggerMapMutex.Lock()
	defer loggerMapMutex.Unlock()

	for _, cl := range loggerMap {
		cl.lvl = cl.calcLevel()
	}
}
