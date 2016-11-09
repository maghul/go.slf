package slf

import (
	"flag"
	"fmt"
	"strings"
	//	"unsafe"
)

type loggerFlagValue struct {
	l *Logger
}

func (lfv loggerFlagValue) String() string {
	if lfv.l == nil { // The undocumented "Zero" Value
		return "Off"
	}
	//	lfv.Stats("String")
	return lfv.l.Level().String()
	return "Off"
}

func (lfv loggerFlagValue) Set(level string) error {
	lvl, err := FindLevel(level)
	if err != nil {
		return err
	}
	lfv.l.SetLevel(lvl)
	return nil
}

func (lfv loggerFlagValue) Stats(prefix string) {
	//	fmt.Println(prefix, " lfv=", unsafe.Pointer(lfv))
	fmt.Println(prefix, " lfv.l=", lfv.l)
	if lfv.l == nil {
		panic("Boom")
		return
	}
	fmt.Println(prefix, " lfv.l.Level()=", lfv.l.Level())

}

func (lfv loggerFlagValue) getUsage() string {
	l := lfv.l
	return fmt.Sprintf("Set the level of logging for '%s', legal values are 'Off', 'Info', and 'Debug', currently set to '%s'",
		l.Description(), l.Level().String())

}
func (l *Logger) initFlag() {
	lfv := loggerFlagValue{l}
	name := fmt.Sprint("slf.", l.name)
	flag.Var(lfv, name, lfv.getUsage())
	//	lfv.Stats("init")

}

func (l *Logger) updateFlagUseage() {
	flag.VisitAll(func(f *flag.Flag) {
		lfv, ok := f.Value.(loggerFlagValue)
		if ok && lfv.l == l {
			f.Usage = lfv.getUsage()
		}
	})
}
