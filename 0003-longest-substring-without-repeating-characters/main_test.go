package main

import (
	"fmt"
	"testing"
)

var tests = []struct {
	Input  string
	Length int
}{
	{Input: "bbbbb", Length: 1},
	{Input: "abcabcbb", Length: 3},
	{Input: "pwwkew", Length: 3},
}

func TestSolutionV0(t *testing.T) {
	for _, v := range tests {
		result := SolutionV0(v.Input)
		Assert(t, v.Length == result, v.Input)
	}
}

func TestSolutionV1(t *testing.T) {
	for _, v := range tests {
		result := SolutionV1(v.Input)
		Assert(t, v.Length == result, v.Input)
	}
}

func TestSolutionV2(t *testing.T) {
	for _, v := range tests {
		result := SolutionV2(v.Input)
		Assert(t, v.Length == result, v.Input)
	}
}

func Assert(tb testing.TB, condition bool, a ...interface{}) {
	tb.Helper()
	if !condition {
		if msg := fmt.Sprint(a...); msg != "" {
			tb.Fatal("Assert failed: " + msg)
		} else {
			tb.Fatal("Assert failed")
		}
	}
}

func Assertf(tb testing.TB, condition bool, format string, a ...interface{}) {
	tb.Helper()
	if !condition {
		if msg := fmt.Sprintf(format, a...); msg != "" {
			tb.Fatal("Assert failed: " + msg)
		} else {
			tb.Fatal("Assert failed")
		}
	}
}
