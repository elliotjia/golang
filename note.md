## 数组和切片
`注意` 绝对不要用指针指向 slice。切片本身已经是一个引用类型，所以它本身就是一个指针!

```
var slice1 []type = arr1[:]

slice1 = &arr1
```

想去掉 slice1 的最后一个元素，只要 `slice1 = slice1[:len(slice1)-1]`
