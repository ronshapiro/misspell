package misspell

import (
	"strings"
	"testing"
)

func TestReplaceIgnore(t *testing.T) {
	cases := []struct {
		ignore string
		text   string
	}{
		{"knwo,gae", "https://github.com/Unknwon, github.com/hnakamur/gaesessions"},
	}

	for line, tt := range cases {
		Ignore(strings.Split(tt.ignore, ","))
		got := ReplaceDebug(tt.text)
		if got != tt.text {
			t.Errorf("%d: Replace files want %q got %q", line, tt.text, got)
		}
	}
}

func TestReplace(t *testing.T) {
	cases := []struct {
		orig string
		want string
	}{
		{"I live in Amercia", "I live in America"},
		{"There is a zeebra", "There is a zebra"},
		{"foo other bar", "foo other bar"},
		{"ten fiels", "ten fields"},
		{"Closeing Time", "Closing Time"},
		{"closeing Time", "closing Time"},
	}
	for line, tt := range cases {
		got := Replace(tt.orig)
		if got != tt.want {
			t.Errorf("%d: Replace files want %q got %q", line, tt.orig, got)
		}
	}
}

func TestReplaceGo(t *testing.T) {
	cases := []struct {
		orig string
		want string
	}{
		{
			orig: `
// I am a zeebra
var foo 10
`,
			want: `
// I am a zebra
var foo 10
`,
		},
		{
			orig: `
var foo 10
// I am a zeebra`,
			want: `
var foo 10
// I am a zebra`,
		},
		{
			orig: `
// I am a zeebra
var foo int
/* multiline
 * zeebra
 */
`,
			want: `
// I am a zebra
var foo int
/* multiline
 * zebra
 */
`,
		},
	}

	for casenum, tt := range cases {
		got := ReplaceGo(tt.orig, true)
		if got != tt.want {
			t.Errorf("%d: %q got converted to %q", casenum, tt, got)
		}
	}
}

func TestCommonPrefixWordLength(t *testing.T) {
	cases := []struct {
		a   string
		b   string
		col int
	}{
		{"", "", 0},
		{"1", "1", 1},
		{"11", "11", 2},
		{"11", "22", 0},
		{"1", "22", 0},
		{"22", "1", 0},
		{"1", "11", 1},
		{"11", "1", 1},
	}

	for casenum, tt := range cases {
		col := commonPrefixWordLength(tt.a, tt.b)
		if col != tt.col {
			t.Errorf("%d: with %q, %q want prefix length of %d, got %d", casenum, tt.a, tt.b, tt.col, col)
		}
	}
}