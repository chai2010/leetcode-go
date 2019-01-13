# 1. 两数之和

给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。

## 示例

给定 `nums = [2, 7, 11, 15], target = 9`

因为 `nums[0] + nums[1] = 2 + 7 = 9`

所以返回 `[0, 1]`

## 方案V0（44ms）

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

1. 代码的实现够简单，采用了嵌套的两次迭代
1. 充分利用了`range`和切片的特性，在内部循环避免了重复判断
1. 时间复杂度是`O(n*n)`

## 方案V1（8ms）

```go
func twoSumV1(nums []int, target int) []int {
	var idxMap = map[int][]int{}
	for i, v := range nums {
		idxMap[v] = append(idxMap[v], i)
	}

	for i, v := range nums {
		if v+v == target {
			if len(idxMap[v]) == 2 {
				return idxMap[v]
			}
			continue
		}

		if _, ok := idxMap[target-v]; ok {
			return []int{i, idxMap[target-v][0]}
		}
	}

	return nil
}
```

1. 采用非嵌套的两次循环，第一次为每个数所在的下标建立索引
1. 第二次遍历的时候，通过map中是否存在需要的数
1. 时间复杂度是`O(n+n+k)`，其中k是map的时间消耗

## 方案V2（8ms）

```go
func twoSumV2(nums []int, target int) []int {
	var idxMap = make(map[int]int32)
	for i, v := range nums {
		idxMap[v] = int32(i)
	}

	for i, v := range nums {
		if j, ok := idxMap[target-v]; ok && int(j) > i {
			return []int{i, int(j)}
		}
	}

	return nil
}
```

在V1中，map的索引不够紧凑。根据题目“你可以假设每种输入只会对应一个答案”，每个数字只会存在一次，因此只需要保存1个索引。从工程角度看，在64位系统，这种方式每个map元素只需要1个`int`，而之前的方式切片头部加元素至少需要4个`int`。更进一步，在Go语言中切片的总大小不超过4GB，因此索引只需要`int32`类型就可以保存了。

V2和V1的数据结构是一样的，因此有着相同的时间复杂度。但是V2的工程实现细节比V1更好。