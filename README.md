# leetcode-go

[![Build Status](https://travis-ci.org/chai2010/leetcode-go.svg)](https://travis-ci.org/chai2010/leetcode-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/chai2010/leetcode-go)](https://goreportcard.com/report/github.com/chai2010/leetcode-go)
[![License](http://img.shields.io/badge/license-BSD-blue.svg)](https://github.com/chai2010/leetcode-go/blob/master/LICENSE)

算法之神Donald Knuth说过：过早优化是万恶之源。Go语言的哲学是：Less is more。根据Go语言哲学我们可以推导出以下推论：

- Slow is fast：代码写的慢一点，代码执行可以变得快一点
- Slow is fast：如果程序跑的慢一点，代码可以写的快一点
- Fast is slow：代码写的很快，那么运行就很容易变慢，同时会拖累项目开发进度变慢
- Fast is slow：如果程序跑的快，写代码就很容易变慢
- 百米赛跑的冠军注定赢不了马拉松的冠军

因此，在满足要求的前提下，我更喜欢符合Go语言哲学的方案。

https://leetcode-cn.com/chai2010/

## 辅助函数

为了便于测试，构造了`Assert`和`Assertf`函数：

```go
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
```

在单元测试中可以这样使用：

```go
func TestFoo(*testing.T) {
	Assert(t, 1+1 == 2, "some", "message")
}
func TestBar(*testing.T) {
	Assertf(t, 1+1 == 2, "%s, %s", "some", "message")
}
```

## BUGS

Report bugs to <chaishushan@gmail.com>.

Thanks!
