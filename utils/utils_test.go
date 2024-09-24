package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNormalizeNewlinesWindows(t *testing.T) {
	testValue := "something\r\nsomething else"
	expectedValue := "something\nsomething else"

	result := NormalizeNewlines([]byte(testValue))

	assert.Equal(t, result, []byte(expectedValue))
}
func TestNormalizeNewlinesMac(t *testing.T) {
	testValue := "something\rsomething else"
	expectedValue := "something\nsomething else"

	result := NormalizeNewlines([]byte(testValue))

	assert.Equal(t, result, []byte(expectedValue))
}

func TestConvertToOsNewLines(t *testing.T) {
	testValue := "something\nsomething else"
	expectedValue := "something\r\nsomething else"

	result := ConvertToOsNewLines(testValue)

	assert.Equal(t, result, expectedValue)
}
