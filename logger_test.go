package slf

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogger(t *testing.T) {

	l1 := GetLogger("test.test")
	l2 := GetLogger("test.test")
	assert.Equal(t, l1, l2)
	assert.Equal(t, "test.test", l1.Name())
}

func TestLoggerOut(t *testing.T) {
	b := bytes.NewBufferString("")
	l1 := GetLogger("test.test")
	l1.SetOutputLogger(b)
	l1.SetLevel(Info)

	l1.Info.Println("Test", 1, 'c')
	assert.Equal(t, "Test 1 99\n", b.String())

	l1.Info.Println("Test", 2)
	assert.Equal(t, "Test 1 99\nTest 2\n", b.String())

	l1.Debug.Println("Test", 3)
	assert.Equal(t, "Test 1 99\nTest 2\n", b.String())

	l1.SetLevel(Debug)

	l1.Debug.Println("Test", 4)
	assert.Equal(t, "Test 1 99\nTest 2\nTest 4\n", b.String())
}
