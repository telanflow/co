package co

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReplace(t *testing.T) {
	asserts := assert.New(t)

	asserts.Equal("origin", Replace(map[string]string{
		"123": "321",
	}, "origin"))

	asserts.Equal("321origin321", Replace(map[string]string{
		"123": "321",
	}, "123origin123"))
	asserts.Equal("321new321", Replace(map[string]string{
		"123":    "321",
		"origin": "new",
	}, "123origin123"))
}

func TestSplitStr(t *testing.T) {
	ss := SplitStr("a, , b,c", ",")
	assert.Equal(t, `[]string{"a", "b", "c"}`, fmt.Sprintf("%#v", ss))

	ss = SplitStr(" ", ",")
	assert.Nil(t, ss)
}

func TestSubStr(t *testing.T) {
	assert.Equal(t, "abc", SubStr("abcDef", 0, 3))
	assert.Equal(t, "cD", SubStr("abcDef", 2, 2))
	assert.Equal(t, "", SubStr("abcDEF", 23, 5))
}

func TestRepeatStr(t *testing.T) {
	assert.Equal(t, "aaa", RepeatStr("a", 3))
	assert.Equal(t, "D", RepeatStr("D", 1))
	assert.Equal(t, "D", RepeatStr("D", 0))
	assert.Equal(t, "D", RepeatStr("D", -3))
}

func TestRepeatRune(t *testing.T) {
	tests := []struct {
		want  []rune
		give  rune
		times int
	}{
		{[]rune("bbb"), 'b', 3},
		{[]rune("..."), '.', 3},
		{[]rune("  "), ' ', 2},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.want, RepeatRune(tt.give, tt.times))
	}
}

func TestPadding(t *testing.T) {
	tests := []struct {
		want, give, pad string
		len             int
		pos             uint8
	}{
		{"ab000", "ab", "0", 5, PosRight},
		{"000ab", "ab", "0", 5, PosLeft},
		{"ab012", "ab012", "0", 4, PosLeft},
		{"ab   ", "ab", "", 5, PosRight},
		{"   ab", "ab", "", 5, PosLeft},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.want, Padding(tt.give, tt.pad, tt.len, tt.pos))

		if tt.pos == PosRight {
			assert.Equal(t, tt.want, PadRight(tt.give, tt.pad, tt.len))
		} else {
			assert.Equal(t, tt.want, PadLeft(tt.give, tt.pad, tt.len))
		}
	}
}
