package main

import (
	"fmt"
	"strings"
)

func main() {
	for _, s := range []string{"bbbbb", "abcabcbb", "pwwkew"} {
		fmt.Println(s, lengthOfLongestSubstring(s))
	}

	// Output:
	// bbbbb 1
	// abcabcbb 3
	// pwwkew 3
}

func lengthOfLongestSubstring(s string) int {
	return SolutionV2(s)
}

func SolutionV0(s string) int {
	if s == "" {
		return 0
	}

	var count = make([]int, len(s))
	for i := 0; i < len(s); i++ {
		count[i] = SolutionV0_strUniqueLen(s[i:])
	}

	var max = count[0]
	for i := 0; i < len(s); i++ {
		if max < count[i] {
			max = count[i]
		}
	}
	return max
}

func SolutionV0_strUniqueLen(s string) int {
	for i := 1; i < len(s); i++ {
		if strings.IndexByte(s[:i], byte(s[i])) >= 0 {
			return i
		}
	}
	return len(s)
}

func SolutionV1(s string) int {
	if s == "" {
		return 0
	}

	var count = make([]int, len(s))
	for i := 0; i < len(s); i++ {
		count[i] = SolutionV1_strUniqueLen(s[i:])
	}

	var max = count[0]
	for i := 0; i < len(s); i++ {
		if max < count[i] {
			max = count[i]
		}
	}
	return max
}

func SolutionV1_strUniqueLen(s string) int {
	var flagMap [256]bool
	for i := 0; i < len(s); i++ {
		if flagMap[s[i]] {
			return i
		}
		flagMap[s[i]] = true
	}
	return len(s)
}

func SolutionV2(s string) int {
	if s == "" {
		return 0
	}

	var max = 0
	for i := 0; i < len(s); i++ {
		count := SolutionV2_strUniqueLen(s[i:])
		if max < count {
			max = count
		}
	}
	return max
}

type FlagMap256 [4]uint64

func (p *FlagMap256) Get(i byte) bool {
	idx0, idx1 := i/64, i%64
	return ((*p)[idx0] & (1 << uint(idx1))) != 0
}
func (p *FlagMap256) SetTrue(i byte) {
	idx0, idx1 := i/64, i%64
	(*p)[idx0] |= (1 << uint(idx1))
}

func SolutionV2_strUniqueLen(s string) int {
	var flagMap FlagMap256
	for i := 0; i < len(s); i++ {
		if flagMap.Get(s[i]) {
			return i
		}
		flagMap.SetTrue(s[i])
	}
	return len(s)
}
