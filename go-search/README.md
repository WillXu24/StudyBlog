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

