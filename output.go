package slf

import (
	"errors"
	"fmt"
)

/*
This is the internal interface used for logging output. It can
be used directly in SetOutputLogger to provide a custom formatting
and logging.
*/
type Output interface {
	Print(ref string, lvl Level, d ...interface{})
	Printf(ref string, lvl Level, args string, d ...interface{})
}

/*
Set an output for the Logger. Will (or should) accept
most common logger frameworks. It will accept the slf.Output
interface, an io.Writer, a core go logger as well as logrus logger.
*/
func (l *Logger) SetOutputLogger(target interface{}) error {
	if target == nil {
		l.own_output = false
		l.out = nil
	} else {
		l.own_output = true
		ad := getAdapter(target)
		if ad == nil {
			return errors.New(fmt.Sprint("Could not find adapter for ", target))
		}
		l.out = ad
	}
	scanSetOutputLoggers()
	return nil
}

func scanSetOutputLoggers() {
	loggerMapMutex.Lock()
	defer loggerMapMutex.Unlock()

	for _, cl := range loggerMap {
		cl.out = cl.calcOutput()
	}
}

// Calculate the level based on parent levels.
func (l *Logger) calcOutput() Output {
	for l != nil {
		if l.own_output {
			return l.out
		}
		l = l.parent
	}
	return output_nilAdapter
}
