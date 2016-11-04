package slf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLevelNames(t *testing.T) {
	assert.Equal(t, "Off", Off.String())
	assert.Equal(t, "Info", Info.String())
	assert.Equal(t, "Parent", Parent.String())
	assert.Equal(t, "Unknown", Level(-4711).String())
}

func checkFindLevel(t *testing.T, expected Level, name string) {
	lvl, err := FindLevel(name)
	assert.NoError(t, err)
	assert.Equal(t, expected, lvl)
}
func TestFindLevel(t *testing.T) {
	checkFindLevel(t, Off, "off")
	checkFindLevel(t, Debug, "debug")
	checkFindLevel(t, Info, "info")
	checkFindLevel(t, Parent, "parent")
	_, err := FindLevel("xyzzy")
	assert.Error(t, err)
}

func TestSetLevels(t *testing.T) {
	l := GetLogger("x")
	assert.Equal(t, "Logger:x, level=Off, output=Output{}", l.String())
	err := SetLevel("x", "Debug")
	assert.NoError(t, err)
	assert.Equal(t, "Logger:x, level=Debug, output=Output{}", l.String())

	err = SetLevel("y", "Debug")
	assert.Error(t, err)

	err = SetLevel("x", "Xyzzy")
	assert.Error(t, err)
}

func TestParentLevel(t *testing.T) {
	p := GetLogger("parent")
	c := GetLogger("child")
	s := GetLogger("subchild")

	c.SetParent(p)
	s.SetParent(c)

	assert.Equal(t, p, c.Parent())
	assert.Equal(t, c, s.Parent())

	assert.Equal(t, "Logger:parent, level=Off, output=Output{}", p.String())
	assert.Equal(t, "Logger:child, level=Off, output=Output{}", c.String())
	assert.Equal(t, "Logger:subchild, level=Off, output=Output{}", s.String())
	p.SetLevel(Debug)
	assert.Equal(t, "Logger:parent, level=Debug, output=Output{}", p.String())
	assert.Equal(t, "Logger:child, level=Debug, output=Output{}", c.String())
	assert.Equal(t, "Logger:subchild, level=Debug, output=Output{}", s.String())
	c.SetLevel(Info)
	assert.Equal(t, "Logger:parent, level=Debug, output=Output{}", p.String())
	assert.Equal(t, "Logger:child, level=Info, output=Output{}", c.String())
	assert.Equal(t, "Logger:subchild, level=Info, output=Output{}", s.String())
	p.SetLevel(Off)
	assert.Equal(t, "Logger:parent, level=Off, output=Output{}", p.String())
	assert.Equal(t, "Logger:child, level=Info, output=Output{}", c.String())
	assert.Equal(t, "Logger:subchild, level=Info, output=Output{}", s.String())
	c.SetLevel(Parent)
	assert.Equal(t, "Logger:parent, level=Off, output=Output{}", p.String())
	assert.Equal(t, "Logger:child, level=Off, output=Output{}", c.String())
	assert.Equal(t, "Logger:subchild, level=Off, output=Output{}", s.String())
}
