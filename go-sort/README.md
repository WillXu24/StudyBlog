## go-sort

> 常见排序算法的golang实现

### 冒泡排序法

```go
func bubbleSort(a []int) {
    // 逆序列对存在标志
	flag := true
    
    // 从头开始遍历
	for i := 0; i < len(a) && flag; i++ {
        // 假设不存在逆序列
		flag = false
        // 从尾部开始进行每一对两数序列比对
		for j := len(a) - 1; j > i; j-- {
            // 若逆序，则调换顺序
			if a[j-1] > a[j] {
				a[j-1], a[j] = a[j], a[j-1]
                // 若存在交换，即存在逆序列对
				flag = true
			}
		}
	}
	fmt.Println("Bubble Sort:", a)
}
```



### 选择排序法

```go
func selectSort(a []int) {
    // 从头开始遍历
	for i := 0; i < len(a); i++ {
        // 设定最小下标
		min := i
        // 从后面开始进行对比
		for j := i + 1; j < len(a); j++ {
            // 如果存在更小的值，则更新最小下标
			if a[min] > a[j] {
				min = j
			}
		}
        // 如果最小下标改变，则调换
		if min != i {
			a[i], a[min] = a[min], a[i]
		}
	}
	fmt.Println("Select Sort:", a)
}
```



### 插入排序法

```go
func insertSort(a []int) {
    // 从头开始遍历
	for i := 1; i < len(a); i++ {
        // 取出一个数
		temp := a[i]
        // 声明该数最后的放置位置
		position := i
        // 从该数的位置往左找的符合条件的位置
		for j := i; j > 0 && a[j-1] > temp; j-- {
            // 每次寻找都将不符合的数右移一位
			a[j] = a[j-1]
			position = j - 1
		}
		a[position] = temp
	}
	fmt.Println("Insert Sort:", a)
}
```



### 希尔排序法

```go
func shellSort(a []int) {
    // 定义间隔，采用2的次方缩小
	gap := 0
	for ; gap < len(a); gap = gap*2 + 1 {
	}
    // 间隔从大到小，直至为1
	for gap = (gap - 1) / 2; gap > 0; gap = (gap - 1) / 2 {
        // 从间隔值开始遍历，方式同插入排序
		for i := gap; i < len(a); i++ {
			temp := a[i]
			position := i
            // 注意此处循环命令不同
			for j := i; j-gap >= 0 && a[j-gap] > temp; j = j - gap {
				a[j] = a[j-gap]
				position = j - gap
			}
			a[position] = temp
		}
	}
	fmt.Println("Shell Sort:", a)
}
```



### 堆排序

```go
func heapSort(a []int) {
    // 遍历所有非叶子节点，从最接近叶子节点的非叶子节点遍历到根节点
	for i := len(a) / 2; i >= 0; i-- {
        // 堆排序
		heapAdjust(a, i)
	}
    // 从倒数第二个节点开始循环遍历
	for i := len(a) - 1; i > 0; i-- {
        // 将最大值调换到序列尾部
		a[0], a[i] = a[i], a[0]
        // 堆排序
		heapAdjust(a[:i], 0)
	}
	fmt.Println("Heap Sort:", a)
}

// 堆排序：将最大数调换到根节点
func heapAdjust(a []int, root int) {
    // 顺着节点找到最小的位置（不太好表达
	for i := 2*root + 1; i < len(a); i = i*2 + 1 {
        // 如果右子树存在，找出左右子树最大值下标
		if i+1 < len(a) && a[i] < a[i+1] {
			i++
		}
        // 如果根节点大于子树，跳出循环
		if a[root] > a[i] {
			break
		}
        // 根节点调换为最值
		a[root], a[i] = a[i], a[root]
        // 原本的根节点变为i，进行下一次比较，直到调换到最小的位置
		root = i
}
```



### 归并排序

```go
func mergeSort0(a []int) {
	res := mSort(a)
	fmt.Println("Merge Sort:", res)
}

// 将a序列进行排序
func mSort(a []int) []int {
    // 递归退出条件
	if len(a) < 2 {
		return a
	}
    // 二分序列
	m := len(a) / 2
    // 左右半序列分别递归排序
	l := mSort(a[:m])
	r := mSort(a[m:])
    // 将左右序列归并
	return merge(l, r)
}

func merge(a, b []int) []int {
	var res []int
	i, j := 0, 0
    // 双指针序列对比拷贝到新序列
	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			res = append(res, a[i])
			i++
		} else {
			res = append(res, b[j])
			j++
		}
	}
    // 将最后剩余部分拷贝到新序列结尾，这里有一个序列其实已经到了结尾，实际上只拷贝了一个序列
	res = append(res, a[i:]...)
	res = append(res, b[j:]...)
	return res
}
```



### 快速排序

```go
func quickSort(a []int) {
	qSort(a, 0, len(a)-1)
	fmt.Println("Quick Sort:", a)
}

// 对a序列快速排序
func qSort(a []int, left, right int) {
    // 递归退出条件
	if left >= right {
		return
	}
    // 寻找并确定枢轴位置
	p := partition(a, left, right)
    // 对枢轴两边的序列分别进行快速排序
	qSort(a, left, p-1)
	qSort(a, p+1, right)
}

// 寻找并确定枢轴位置
func partition(a []int, left, right int) int {
    // 选择首位为枢轴
	p := a[left]
    // 双指针逼近，重合时退出
	for left != right {
        // 最右值大于枢轴，右指针左移
		for left < right && p < a[right] {
			right--
		}
        // 最右值小于枢轴，调换位置
		a[left], a[right] = a[right], a[left]
        // 最左值小于枢轴，左指针右移
		for left < right && a[left] < p {
			left++
		}
        // 最左值大于枢轴，调换位置
		a[left], a[right] = a[right], a[left]
	}
    // 返回最后指针（枢轴）位置
	return left
}
```

