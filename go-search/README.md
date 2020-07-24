## go-search

> 常见查找算法的golang实现

### 顺序查找

```go
func SequenceSearch(a []int,b int) int {
	for i:=0;i<len(a);i++{
		if a[i]==b{
			return i
		}
	}
	return -1
}
```

算法分析：简单遍历，线性表就可以

复杂度分析：

- 时间复杂度：**O(n)**



### 二分查找

```go
func BinarySearch0(a []int, b int) int {
	l,m,h:=0,0,len(a)-1
	for l<=h{
		m=(h+l)/2
		if a[m]==b{
			return m
		}
		if a[m]>b{
			h=m-1
		}
		if a[m]<b{
			l=m+1
		}
	}
	return -1
}

func BinarySearch1(a []int, b int) int {
return bSearch(a,0,len(a)-1,b)
}

func bSearch(a[]int,l,h,b int) int {
	m:=(l+h)/2
	if a[m]==b{
		return m
	}
	if a[m]>b{
		return bSearch(a,l,m-1,b)
	}
	if a[m]<b{
		return bSearch(a,m+1,h,b)
	}
	return -1
}
```

算法分析：**有序**查找！注意递归实现的思想。

复杂度分析：

- 时间复杂度：**O(log****2****n)**



### 插值查找

```go
func InsertSearch(a []int, b int) int {
	l,m,h:=0,0,len(a)-1
	for l<=h{
		m=l+(b-a[l])/(a[h]-a[l])*(h-l)
		if a[m]==b{
			return m
		}
		if a[m]>b{
			h=m-1
		}
		if a[m]<b{
			l=m+1
		}
	}
	return -1
}
```

算法分析：二分查找的改进版，m值的确定变为自适应，可以在一定程度上提高查找效率。

复杂度分析：

- 时间复杂度：**O(log2(log2n))**



### 斐波那契查找

```go
func FibonacciSearch(a []int ,b int) int {
    // 先构造一个斐波那契数列
	var f []int
	f=append(f,1,1)
	for i:=2;i<100;i++{
		f=append(f,f[i-1]+f[i-2])
	}

    // 计算临时数列的容量
	var n int
	for f[n]-1<len(a){
		n++
	}

    // 创建临时数列
	temp:=make([]int,f[n]-1)
	for i,v:=range a{
		temp[i]=v
	}
    // 后面的部分用最后一位补齐
	for i:=len(a);i<len(temp);i++{
		temp[i]=a[len(a)-1]
	}

    // 查找部分
	l,m,h:=0,0,len(a)
	for l<=h{
        // 黄金分割点
		m = l+f[n-1]-1
		if temp[m] == b {
            // 如果查找到补齐部分返回数列末尾下标
			if m<len(a){
				return m
			}else {
				return len(a)-1
			}
		}
		if temp[m] > b {
			h=m-1
			n=n-1
		}
		if temp[m] < b {
			l = m + 1
			n=n-2
		}
	}

	return -1
}
```

算法分析：二分查找的改进版，利用黄金分割的特点，但是对于数组的长度有要求。

复杂度分析：

- 时间复杂度：**O(log2****n)**