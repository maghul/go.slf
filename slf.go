/*
The slf package is a simple logging facade for Go. It implements
a logger which can be redirected to other logging frameworks and
stdout/stderr. Logging entities can be retreived and modifed between
packages. This is mainly useful for the main package to initialize
and configure logging and library packages can output to the loggers
without worrying about where the logging is sent.
*/
package slf

import (
	"errors"
	"fmt"
	"sort"
	"strings"
	"sync"
)

var loggerMap = make(map[string]*Logger)
var loggerMapMutex sync.Mutex

/*
Create a new Logger. Returns error if already exists.
*/
func CreateLogger(name string) (*Logger, error) {
	n := strings.ToLower(name)

	loggerMapMutex.Lock()
	defer loggerMapMutex.Unlock()

	if _, ok := loggerMap[n]; ok {
		return nil, errors.New(fmt.Sprintf("Could not create logger '%s', it already exists", name))
	}
	logger := makeLogger(n)
	loggerMap[n] = logger
	return logger, nil
}

/*
Get a Logger. Returns an existing logger or creates
a new logger.
*/
func GetLogger(name string) *Logger {
	n := strings.ToLower(name)

	loggerMapMutex.Lock()
	defer loggerMapMutex.Unlock()

	if _, ok := loggerMap[n]; ok {
		return loggerMap[n]
	}
	logger := makeLogger(n)
	loggerMap[n] = logger
	return logger
}

/*
Find a logger. Returns an existing Logger or error
if it doesn't exist.
*/
func FindLogger(name string) (*Logger, error) {
	n := strings.ToLower(name)

	loggerMapMutex.Lock()
	defer loggerMapMutex.Unlock()

	if _, ok := loggerMap[n]; ok {
		return loggerMap[n], nil
	}
	return nil, errors.New(fmt.Sprintf("Could not find logger '%s', it hasn't been created", name))
}

/*
Delete a Logger if it exists. Will return an error if it doesn't
*/
func (l *Logger) Delete() error {
	loggerMapMutex.Lock()
	defer loggerMapMutex.Unlock()

	n := l.name

	cl, ok := loggerMap[n]
	if ok {
		if cl != l {
			return errors.New(fmt.Sprintf("Could not delete logger '%s', it is duplicated or mismatched"))
		}
		delete(loggerMap, n)
		return nil
	}
	return errors.New(fmt.Sprintf("Could not delete logger '%s', it hasn't been created"))

}

type loggers []*Logger

/*
Returns all existing loggers. Will be sorted alphabetically by name.
*/
func Loggers() []*Logger {
	loggerMapMutex.Lock()
	defer loggerMapMutex.Unlock()

	rv := make([]*Logger, len(loggerMap))
	ii := 0
	for _, v := range loggerMap {
		rv[ii] = v
		ii++
		fmt.Println("Adding logger ", ii, ", logger=", v)
	}

	if len(rv) > 0 {
		sort.Sort(loggers(rv))
	}
	return rv
}

func (ls loggers) Len() int {
	fmt.Println("Len:", ls)
	return len(ls)
}

func (ls loggers) Less(i, j int) bool {
	fmt.Println("Less:", i, ", ")
	fmt.Println("Less:", j, ", ")
	return strings.ToLower(ls[i].Name()) < strings.ToLower(ls[j].Name())
}

func (ls loggers) Swap(i, j int) {
	x := ls[i]
	ls[i] = ls[j]
	ls[j] = x
}
