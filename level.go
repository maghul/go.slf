package slf

import (
	"errors"
	"fmt"
	"strings"
)

type Level int

const (
	Off Level = iota
	Info
	Debug
	Parent Level = -1
)

/* Output the level name as a string.
 */
func (lvl Level) String() string {
	switch lvl {
	case Off:
		return "Off"
	case Info:
		return "Info"
	case Debug:
		return "Debug"
	case Parent:
		return "Parent"
	default:
		return "Unknown"
	}
}

/*
Get the level by name. The levels are 'info', 'debug', 'parent'and 'off'
Returns an error if the level can't be found.
*/
func FindLevel(level string) (Level, error) {
	switch strings.ToLower(level) {
	case "off":
		return Off, nil
	case "info":
		return Info, nil
	case "debug":
		return Debug, nil
	case "parent":
		return Parent, nil
	default:
		return Off, errors.New(fmt.Sprint("There is no level named '", level, "'"))
	}
}
