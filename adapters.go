package slf

import (
	"fmt"
	"io"
)

// Implements and adapter for Println style loggers.
type printlnlogger interface {
	Println(d ...interface{}) (int, error)
}

type output_printlnlogger struct {
	w printlnlogger
}

func (o *output_printlnlogger) Print(ref string, lvl Level, d ...interface{}) {
	o.w.Println(d...)
}

func (o *output_printlnlogger) Printf(ref string, lvl Level, arg string, d ...interface{}) {
	o.w.Println(fmt.Sprintf(arg, d...))
}

func (o *output_printlnlogger) String() string {
	return "Output{println}"
}

// Implements an adapter for standard Go Logger.
type logLogger interface {
	Output(calldepth int, s string) error
}

type output_logLogger struct {
	w logLogger
}

func (o *output_logLogger) Print(ref string, lvl Level, d ...interface{}) {
	o.w.Output(2, fmt.Sprintln(d...))
}

func (o *output_logLogger) Printf(ref string, lvl Level, args string, d ...interface{}) {
	o.w.Output(2, fmt.Sprintf(args, d...))
}

func (o *output_logLogger) String() string {
	return "Output{Logger}"
}

// implements io.Writer which is used by lumberjack and stdio
type output_ioWriter struct {
	w io.Writer
}

func (o *output_ioWriter) Print(ref string, lvl Level, d ...interface{}) {
	fmt.Fprintln(o.w, d...)
}

func (o *output_ioWriter) Printf(ref string, lvl Level, args string, d ...interface{}) {
	fmt.Fprintf(o.w, args, d...)
}

func (o *output_ioWriter) String() string {
	return "Output{ioWriter}"
}

// Implements NIL writer.
type output_nil struct {
}

func (o *output_nil) Print(ref string, lvl Level, d ...interface{}) {
}

func (o *output_nil) Printf(ref string, lvl Level, args string, d ...interface{}) {
}

func (o *output_nil) String() string {
	return "Output{}"
}

var output_nilAdapter = &output_nil{}

// --------------------------------------------------------------------------
type testOutput struct {
	ref string
}

func (to *testOutput) Print(ref string, lvl Level, d ...interface{}) {
}

func (to *testOutput) Printf(ref string, lvl Level, arg string, d ...interface{}) {
}

func (to *testOutput) String() string {
	return to.ref
}

// --------------------------------------------------------------------------
func getAdapter(target interface{}) Output {
	if target == nil {
		return output_nilAdapter
	}
	switch t := target.(type) {
	case Output:
		return t
	case io.Writer:
		return &output_ioWriter{t}
	case logLogger:
		return &output_logLogger{t}
	case printlnlogger:
		return &output_printlnlogger{t}
	case *testOutput:
		return t
	default:
	}
	return nil
}
