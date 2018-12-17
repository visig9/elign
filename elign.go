// Package elign align the string with east asian character set.
package elign

import (
	"strings"

	"golang.org/x/text/width"
)

// WidthFn is function can calculate the width size of a string.
type WidthFn func(s string) (width int)

// Width is build-in WidthFn can retrieve the width of a string.
func Width(s string) (w int) {
	for _, r := range []rune(s) {
		switch width.LookupRune(r).Kind() {
		case width.EastAsianWide, width.EastAsianFullwidth:
			w += 2
		default:
			w++
		}
	}

	return
}

// PadSpaceNum retrieve the spaces number should used to pad.
//
// cw is column width.
func padSpaceNum(wfn WidthFn, cw int, s string) int {
	num := wfn(s)

	if delta := cw - num; delta > 0 {
		return delta
	}

	return 0
}

// Spaces generate a string contain n spaces.
func spaces(n int) string {
	return strings.Repeat(" ", n)
}

// Right padding whitespace to the left of string if cw wider than
// the string size.
//
// cw is column width.
func Right(wfn WidthFn, cw int, s string) string {
	return spaces(padSpaceNum(wfn, cw, s)) + s
}

// Left padding whitespace to the right of string if cw wider than
// the string size.
//
// cw is column width.
func Left(wfn WidthFn, cw int, s string) string {
	return s + spaces(padSpaceNum(wfn, cw, s))
}

// Elign is using for align text with east character set.
type Elign struct {
	// WidthFn offer an algorithm can retrive the width of a string.
	WidthFn WidthFn
	// ColumnWidth offer the filed size of align operations.
	ColumnWidth int
}

// Default return a default Elign instance (using Width as WidthFn).
// This function only for convenient.
//
// cw is initial column width.
func Default(cw int) *Elign {
	return &Elign{
		WidthFn:     Width,
		ColumnWidth: cw,
	}
}

// AdjustWidth enlarge the ColumnWidth if it's smaller than input string.
//
// It will return itself for chaining call.
func (e *Elign) AdjustWidth(strs ...string) *Elign {
	for _, s := range strs {
		if w := e.WidthFn(s); w > e.ColumnWidth {
			e.ColumnWidth = w
		}
	}

	return e
}

// Right padding whitespace to the left of string if cw wider than
// the string size.
func (e *Elign) Right(s string) string {
	return Right(e.WidthFn, e.ColumnWidth, s)
}

// Left padding whitespace to the right of string if cw wider than
// the string size.
func (e *Elign) Left(s string) string {
	return Left(e.WidthFn, e.ColumnWidth, s)
}
