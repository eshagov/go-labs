package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFuncUncompress(t *testing.T) {
	var str1 = "e3r4"
	var str2 = "eeerrrr"

	assert.Equal(t, str2, FuncUncompress(str1), "Two string must be equals")
}
