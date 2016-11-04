package slf

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParentOutput(t *testing.T) {
	p := GetLogger("parent")
	c := GetLogger("child")
	s := GetLogger("subchild")

	c.SetParent(p)
	s.SetParent(c)

	assert.Equal(t, "Logger:parent, level=Off, output=Output{}", p.String())
	assert.Equal(t, "Logger:child, level=Off, output=Output{}", c.String())
	assert.Equal(t, "Logger:subchild, level=Off, output=Output{}", s.String())
	p.SetOutputLogger(&testOutput{"A"})
	assert.Equal(t, "Logger:parent, level=Off, output=A", p.String())
	assert.Equal(t, "Logger:child, level=Off, output=A", c.String())
	assert.Equal(t, "Logger:subchild, level=Off, output=A", s.String())
	c.SetOutputLogger(&testOutput{"B"})
	assert.Equal(t, "Logger:parent, level=Off, output=A", p.String())
	assert.Equal(t, "Logger:child, level=Off, output=B", c.String())
	assert.Equal(t, "Logger:subchild, level=Off, output=B", s.String())
	s.SetOutputLogger(&testOutput{"C"})
	assert.Equal(t, "Logger:parent, level=Off, output=A", p.String())
	assert.Equal(t, "Logger:child, level=Off, output=B", c.String())
	assert.Equal(t, "Logger:subchild, level=Off, output=C", s.String())
	s.SetOutputLogger(nil)
	assert.Equal(t, "Logger:parent, level=Off, output=A", p.String())
	assert.Equal(t, "Logger:child, level=Off, output=B", c.String())
	assert.Equal(t, "Logger:subchild, level=Off, output=B", s.String())
	c.SetOutputLogger(nil)
	assert.Equal(t, "Logger:parent, level=Off, output=A", p.String())
	assert.Equal(t, "Logger:child, level=Off, output=A", c.String())
	assert.Equal(t, "Logger:subchild, level=Off, output=A", s.String())
}

type MockDirectOutput struct {
	buf bytes.Buffer
}

func makeMockDirectOutput(t *testing.T) *MockDirectOutput {
	mdo := &MockDirectOutput{}
	assert.Implements(t, (*Output)(nil), mdo, "MockDirectOutput should implement Output")
	return mdo
}

func (mdo *MockDirectOutput) Print(ref string, lvl Level, d ...interface{}) {
	fmt.Fprint(&mdo.buf, ref, ":", lvl, ":")
	fmt.Fprintln(&mdo.buf, d...)
}

func (mdo *MockDirectOutput) Printf(ref string, lvl Level, args string, d ...interface{}) {
	fmt.Fprint(&mdo.buf, ref, ":", lvl, ":")
	fmt.Fprintf(&mdo.buf, args, d...)
}

func (mdo *MockDirectOutput) String() string {
	return "MockDirecttOutput"
}

func (mdo *MockDirectOutput) GetLog() string {
	s := mdo.buf.String()
	mdo.buf.Reset()
	return s
}

func TestDirectOutput(t *testing.T) {
	l := GetLogger("xyz")

	mdo := makeMockDirectOutput(t)
	err := l.SetOutputLogger(mdo)
	assert.NoError(t, err)
	l.SetLevel(Debug)

	l.Debug.Println("message")
	assert.Equal(t, "xyz:Debug:message\n", mdo.GetLog())

	l.Info.Println("hint")
	assert.Equal(t, "xyz:Info:hint\n", mdo.GetLog())

	l.Info.Printf("arg=%d\n", 4711)
	assert.Equal(t, "xyz:Info:arg=4711\n", mdo.GetLog())

}
