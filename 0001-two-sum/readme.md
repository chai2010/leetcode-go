# 1. 两数之和

给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。

## 示例

给定 `nums = [2, 7, 11, 15], target = 9`

因为 `nums[0] + nums[1] = 2 + 7 = 9`

所以返回 `[0, 1]`

## 方案V0

```go
func twoSumV0(nums []int, target int) []int {
	for i, vi := range nums {
		for j, vj := range nums[i+1:] {
			if vi+vj == target {
				return []int{i, i + 1 + j}
			}
		}
	}
	return nil
}
```

这个方案可能不是最快的，但是可能是最符合Go编程哲学的。

1. 代码的实现够简单，采用了两次迭代
1. 充分利用了`range`和切片的特性，在内部循环避免了重复判断
