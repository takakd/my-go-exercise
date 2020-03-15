# My Go Exercise

## Reference

Thanks for publishing good articles.

### Testing
1. test1: https://blog.alexellis.io/golang-writing-unit-tests/

### Aizu Online Judge

[Aizu Online Judge](https://onlinejudge.u-aizu.ac.jp/home)
with [プログラミングコンテスト攻略のためのアルゴリズムとデータ構造](https://www.amazon.co.jp/%E3%83%97%E3%83%AD%E3%82%B0%E3%83%A9%E3%83%9F%E3%83%B3%E3%82%B0%E3%82%B3%E3%83%B3%E3%83%86%E3%82%B9%E3%83%88%E6%94%BB%E7%95%A5%E3%81%AE%E3%81%9F%E3%82%81%E3%81%AE%E3%82%A2%E3%83%AB%E3%82%B4%E3%83%AA%E3%82%BA%E3%83%A0%E3%81%A8%E3%83%87%E3%83%BC%E3%82%BF%E6%A7%8B%E9%80%A0-%E6%B8%A1%E9%83%A8-%E6%9C%89%E9%9A%86/dp/4839952957/ref=sr_1_1?__mk_ja_JP=%E3%82%AB%E3%82%BF%E3%82%AB%E3%83%8A&keywords=%E3%82%A2%E3%83%AB%E3%82%B4%E3%83%AA%E3%82%BA%E3%83%A0+%E7%AB%B6%E6%8A%80&qid=1583866126&s=books&sr=1-1)

## Note

### Struct constructor

```go
type StructName struct {
	value int
}

// constructor
func NewStructName() *StructName {
    s := &StructName()
    // or s := new(StructName)
    
    s.value = 123
    return s
}

// with parameter
func NewStructName(value int) *StructName {
    s := &StructName{value :value}
    return s
}
```

@ref: https://gobyexample.com/structs

### Variable declarations

```go
var i1 int
var i2 int
var i3 int

// are equals

var i1, i2, i3 int
```

@ref: https://golang.org/ref/spec#Variable_declarations
