package slf

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdapter(t *testing.T) {
	x := getAdapter(os.Stdout)

	assert.NotNil(t, x)
}

func TestWriterAdapter(t *testing.T) {
	b := bytes.NewBufferString("")
	x := getAdapter(b)
	assert.NotNil(t, x)
	assert.Equal(t, "Output{ioWriter}", fmt.Sprint(x))

	x.Print("test1", Debug, "a", "b")
	assert.Equal(t, "a b\n", b.String())

	b.Reset()

	x.Printf("test1", Debug, "a'%s'", "b")
	assert.Equal(t, "a'b'", b.String())
}

type printlnAdapter struct {
	b *bytes.Buffer
}

func (pa *printlnAdapter) Println(d ...interface{}) (int, error) {
	fmt.Fprintln(pa.b, d...)
	return 0, nil
}

func TestPrintlnAdapter(t *testing.T) {
	tpa := &printlnAdapter{bytes.NewBufferString("")}
	x := getAdapter(tpa)
	assert.NotNil(t, x)
	assert.Equal(t, "Output{println}", fmt.Sprint(x))

	x.Print("test1", Debug, "a", "b")
	assert.Equal(t, "a b\n", tpa.b.String())

	tpa.b.Reset()

	x.Printf("test1", Debug, "a'%s'", "b")
	assert.Equal(t, "a'b'\n", tpa.b.String())
}

func TestLoggerAdapter(t *testing.T) {
	b := bytes.NewBufferString("")
	logger := log.New(b, "test:", 0)
	x := getAdapter(logger)
	assert.NotNil(t, x)
	assert.Equal(t, "Output{Logger}", fmt.Sprint(x))

	x.Print("test1", Debug, "a", "b")
	assert.Equal(t, "test:a b\n", b.String())

	b.Reset()

	x.Printf("test1", Debug, "a'%s'", "b")
	assert.Equal(t, "test:a'b'\n", b.String())
}

func TestNilAdapter(t *testing.T) {
	x := getAdapter(nil)
	assert.NotNil(t, x)
	assert.Equal(t, "Output{}", fmt.Sprint(x))

	x.Print("test1", Debug, "a", "b")
	x.Printf("test1", Debug, "a'%s'", "b")
}
