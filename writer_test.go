package slf

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWriterOut(t *testing.T) {
	b := bytes.NewBufferString("")
	l1 := GetLogger("test.test")
	l1.SetOutputLogger(b)
	l1.SetLevel(Info)

	wi := l1.Info
	wd := l1.Debug

	fmt.Fprintln(wi, "Test", 1, 'c')
	assert.Equal(t, "Test 1 99\n", b.String())

	fmt.Fprintln(wi, "Test", 2)
	assert.Equal(t, "Test 1 99\nTest 2\n", b.String())

	fmt.Fprintln(wd, "Test", 3)
	assert.Equal(t, "Test 1 99\nTest 2\n", b.String())

	l1.SetLevel(Debug)

	fmt.Fprintln(wd, "Test", 4)
	assert.Equal(t, "Test 1 99\nTest 2\nTest 4\n", b.String())
}
