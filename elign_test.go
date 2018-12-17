package elign

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// fake WidthFn
func wfn(s string) int { return len([]rune(s)) }

func TestWidth(t *testing.T) {
	cases := []struct {
		s string
		w int
	}{
		{"", 0},
		{"123", 3},
		{"一二三", 6},
		{"123一二三", 9},
		{"123，一二三", 11},
	}

	for _, c := range cases {
		assert.Equal(t, c.w, Width(c.s))
	}
}

func TestPadSpaceNum(t *testing.T) {
	cw := 10

	cases := []struct {
		s string
		w int
	}{
		{"", 10},
		{"123", 7},
		{"一二三", 7},
		{"123一二三", 4},
		{"123，一二三", 3},
		{"1234561790123456", 0},
	}

	for _, c := range cases {
		assert.Equal(t, c.w, padSpaceNum(wfn, cw, c.s), "%v", c)
	}
}

func TestSpaces(t *testing.T) {
	cases := []struct {
		n int
		w string
	}{
		{0, ""},
		{1, " "},
		{2, "  "},
		{3, "   "},
	}

	for _, c := range cases {
		assert.Equal(t, c.w, spaces(c.n))
	}
}

func TestRight(t *testing.T) {
	cw := 10

	cases := []struct {
		s string
		w string
	}{
		{"", "          "},
		{"123", "       123"},
		{"一二三", "       一二三"},
		{"123一二三", "    123一二三"},
		{"123，一二三", "   123，一二三"},
		{"1234567890123456", "1234567890123456"},
	}

	for _, c := range cases {
		assert.Equal(t, c.w, Right(wfn, cw, c.s))
	}
}

func TestLeft(t *testing.T) {
	cw := 10

	cases := []struct {
		s string
		w string
	}{
		{"", "          "},
		{"123", "123       "},
		{"一二三", "一二三       "},
		{"123一二三", "123一二三    "},
		{"123，一二三", "123，一二三   "},
		{"1234567890123456", "1234567890123456"},
	}

	for _, c := range cases {
		assert.Equal(t, c.w, Left(wfn, cw, c.s))
	}
}

func TestDefault(t *testing.T) {
	// Flaw & Skip: cannot compare function is identity
	cases := []struct {
		cw int
	}{
		{0},
		{1},
		{5},
		{10},
	}

	for _, c := range cases {
		assert.Equal(t, c.cw, Default(c.cw).ColumnWidth)
	}
}

func TestElign(t *testing.T) {
	cases := []struct {
		ss []string // test data, ss[0] using to run Right & Left
		wr string   // want right
		wl string   // want left
	}{
		{[]string{"", ""}, "", ""},
		{[]string{"", "b"}, " ", " "},
		{[]string{"a", ""}, "a", "a"},
		{[]string{"a", "b"}, "a", "a"},
		{[]string{"a", "bc"}, " a", "a "},
		{[]string{"a", "一"}, " a", "a "},
		{[]string{"a", "一d"}, "  a", "a  "},
		{[]string{"一", "b"}, "一", "一"},
		{[]string{"一", "bc"}, "一", "一"},
		{[]string{"一", "bcd"}, " 一", "一 "},
	}

	for _, c := range cases {
		e := Default(0)
		for _, s := range c.ss {
			e.AdjustWidth(s)
		}

		assert.Equal(t, c.wr, e.Right(c.ss[0]))
		assert.Equal(t, c.wl, e.Left(c.ss[0]))
	}
}

func ExampleElign() {
	data := []string{
		"世界上",
		"只有 10 種人",
		"懂二進位和不懂二進位的",
	}
	e := Default(0).AdjustWidth(data...)

	for _, d := range data {
		fmt.Printf("|%v|\n", e.Right(d))
	}

	// Output:
	// |                世界上|
	// |          只有 10 種人|
	// |懂二進位和不懂二進位的|
}
