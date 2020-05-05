package co

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContainsUint(t *testing.T) {
	asserts := assert.New(t)
	asserts.True(ContainsUint([]uint{0, 2, 3, 65, 4}, 65))
	asserts.True(ContainsUint([]uint{65}, 65))
	asserts.False(ContainsUint([]uint{65}, 6))
}

func TestContainsString(t *testing.T) {
	asserts := assert.New(t)
	asserts.True(ContainsString([]string{"", "1"}, ""))
	asserts.True(ContainsString([]string{"", "1"}, "1"))
	asserts.False(ContainsString([]string{"", "1"}, " "))
}

func TestSliceDifference(t *testing.T) {
	asserts := assert.New(t)

	{
		s1 := []string{"1", "2", "3", "4"}
		s2 := []string{"2", "4"}
		asserts.Equal([]string{"1", "3"}, SliceDifference(s1, s2))
	}

	{
		s2 := []string{"1", "2", "3", "4"}
		s1 := []string{"2", "4"}
		asserts.Equal([]string{}, SliceDifference(s1, s2))
	}

	{
		s1 := []string{"1", "2", "3", "4"}
		s2 := []string{"1", "2", "3", "4"}
		asserts.Equal([]string{}, SliceDifference(s1, s2))
	}

	{
		s1 := []string{"1", "2", "3", "4"}
		s2 := []string{}
		asserts.Equal([]string{"1", "2", "3", "4"}, SliceDifference(s1, s2))
	}
}
