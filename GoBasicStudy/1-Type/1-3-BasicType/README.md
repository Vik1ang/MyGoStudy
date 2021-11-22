# 基本类型

更明确的数字类型命名，⽀持 `Unicode`，⽀持常⽤数据结构

| 类型 | 长度 | 默认值 | 说明 |
| --- | --- | --- | ---|
| `bool` | 1 | false | |
| `byte` | 1 | 0 | `uint8`|
| `rune` | 4 | 0 | `Unicode Code Point`, `int32`|
|`int`, `uint`| 4 或 8 | 0 | 32位 或 64位 |
|`int8`, `uint8`| 1 | 0 | -128 ~ 127, 0 ~ 255 |
|`int16`, `uint16` | 2 | 0 | -32768 ~ 32767, 0 ~ 65535 |
|`int32`, `uint32` | 4 | 0 | -21亿 ~ 21 亿, 0 ~ 42 亿 |
|`int64`, `uint64` | 8 | 0 ||
|`float32`| 4 | 0.0 ||
|`float64`| 8 | 0.0 ||
|`complex64`| 8 | 0 ||
|`complex128`| 16 | 0 ||
|`uintptr`| 4 或 8 | | ⾜以存储指针的 `uint32` 或 `uint64` 整数 |
|`array`| | | 值类型 |
|`struct`| | | 值类型 |
| `string` | | `""` | `UTF-8` 字符串|
| `slice` | | `nil` | 引用类型 |
| `map` | | `nil` | 引用类型 |
| `channel` | | `nil` | 引用类型 |
| `interface` | | `nil` | 接口 |
| `function` | | `nil` | 函数 |

⽀持⼋进制、⼗六进制，以及科学记数法。标准库 math 定义了各数字类型取值范围

```go
a, b, c, d := 071, 0x1F, 1e9, math.MinInt16
```

空指针值 `nil`，⽽⾮ C/C++ NULL。
