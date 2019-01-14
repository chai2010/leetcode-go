# 3. 无重复字符的最长子串

给定一个字符串，找出不含有重复字符的最长子串的长度。
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

## 示例1

1. 输入: `"abcabcbb"`
1. 输出: `3`
1. 解释: 因为无重复字符的最长子串是 `"abc"`，所以其长度为 `3`。

## 示例2

1. 输入: `"bbbbb"`
1. 输出: `1`
1. 解释: 因为无重复字符的最长子串是 `"b"`，所以其长度为 `1`。

## 示例3

1. 输入: `"pwwkew"`
1. 输出: `3`
1. 解释: 因为无重复字符的最长子串是 `"wke"`，所以其长度为 `3`。

请注意，你的答案必须是 子串 的长度，`"pwke"` 是一个子序列，不是子串。

## 方案V0（32ms）

```go
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
	var i = 1
	for ; i < len(s); i++ {
		if strings.IndexByte(s[:i], byte(s[i])) >= 0 {
			return i
		}
	}
	return i
}
```

思路是计算字符串中每个字符开始的新字符串的第一个无重复字符的子串长度，然后统计最大的长度。

## 方案V1（16ms）

```go
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
```

优化`strUniqueLen`函数。假设字符为1个字节，采用`[256]bool`表格标注字符是否出现过。

## 方案V2（16ms）

```go
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
```

优化`strUniqueLen`函数。假设字符为1个字节，采用`[4]uint64`大小的表格标注字符是否出现过。时间复杂度是`O(n*n)`。
