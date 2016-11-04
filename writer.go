package slf

import (
	"strings"
)

/*
Write bytes to the logger. This is an implementation of io.Writer
although expensive. Will split lines as necessary.
*/
func (lo *LoggerLevel) Write(msg []byte) (int, error) {
	if lo.l.lvl < lo.lvl {
		return len(msg), nil
	}
	m := string(msg)
	ms := strings.Split(m, "\n")

	for _, l := range ms {
		ll := len(l)
		if ll > 1 {
			lo.l.print(lo.l.name, lo.lvl, l)
		}
	}
	return len(msg), nil
}
