## 2.2 Go语言基础

### 2.2.1 定义变量

使用var关键字是Go语言最基本的定义变量方式

与C语言不同的是Go语言把变量类型放在变量名后面

```go

//定义一个名称为 "variableName" 类型为"type"的变量
var variableName type;

// 定义多个变量
var vname1, vname2, vname3 type

// 定义变量并初始化值
var variableName type = value

// 同时初始化多个变量
var vname1, vname2, vname3 type = v1, v2, v3

// 简写
var vname1, vname2, vname3 = v1, v2, v3

// 再简写
vname1, vname2, vname3 := v1, v2, v3

```

`:=`这个符号直接取代了`var`和`type`, 这种形式叫做简短声明

不过它有一个**限制**, 那就是它只能用在**函数内部**, 在函数外部使用则会**无法编译通过**，所以一般用var方式来定义全局变量

`_` (下划线) 是个特殊的变量名, 任何赋予它的值都会被丢弃 在这个例子中, 我们将值35赋予b，并同时丢弃34

```go

_, b := 34, 35

```

Go语言对于已声明但**未使用的变量**会在**编译阶段报错**

### 2.2.2 常量

所谓常量, 也就是在程序**编译阶段就确定下来的值**, 而程序在运行时则无法改变该值

在Go语言程序中, 常量可定义为数值,布尔值或字符串等类型

```go

const constantName = value
const  Pi float32 = 3.1415926

const Pi = 3.1415926
const i = 10000
const MaxThread = 10
const prefix = "prefix_"

```

### 2.2.2 内置基础类型

#### Boolean

```go

var isActive bool // 全局变量声明
var enabled, disabled = true, false // 忽略类型的声明

func test() {
	var available bool // 一般声明
	valid := false // 简短声明
	available = true // 赋值操作
}

```

#### 数值类型

整数类型有无符号和带符号两种

Go语言同时支持int和uint, 这两种类型的长度相同, 但具体长度取决于不同编译器的实现

当前的gcc和gccgo编译器在32位和64位平台上都使用32位来表示int和uint, 但未来在64位平台上可能增加到64位

Go语言里面也有直接定义好位数的类型: `rune`, `int8`, `int16`, `int32`, `int64`, `byte`, `uint8`, `uint16`, `uint32`, `uint64`

其中`rune`是`int32`的别称, `byte`是`uint8`的别称

需要注意的一点是, 这些类型的变量之间不允许互相赋值或操作, 不然会在**编译时引起编译器报错**

**尽管int的长度是32 bit，但int与int32并不可以互用**

浮点数的类型有`float32`和`float64`两种(没有float类型), **默认是`float64`**

Go语言还支持复数。它的默认类型是complex128(64位实数＋64位虚数), 如果需要小一些的，也有complex64（32位实数＋32位虚数）

复数的形式为`RE＋IMi`, 其中`RE`是**实数部分**, `IM`是**虚数部分**, 而最后的`i`是**虚数单位**

下面是一个使用复数的例子

```go

var c complex64 = 5 + 5i;
fmt.Printf("Value is: %v", c)

```

#### 字符串

字符串是用一对双引号`""`或反引号(` `)括起来定义, 它的类型是string

```go

var frenchHello string // 声明变量为字符串的一般方法
var emptyString string = "" // 声明一个字符串变量, 初始化为空字符串
func test() {
	no, yes, maybe := "no", "yes", "maybe" // 简短声明, 同时声明多个变量
	japaneseHello := "Ohaiou"
	frenchHello = "Bonjour" // 常规赋值
}

```

在Go语言中字符串是不可变的

以下代码编译时会报错
```go 

var s string = "hello"
s[0] = 'c'

```

但如果真的想要修改怎么办? 下面的代码可以实现

```go

s := "hello"
c := []byte(s) // 将字符串s转换为[]byte类型
c[0] = 'c'
s2 := string(c) // 再转回string类型
fmt.Printf("%s\n", s2)

```

Go语言中可以使用`+`操作符来连接两个字符串

```go

s := "hello"
m := "world"
a := s + m
fmt.Printf("%s\n", a)

```

修改字符串也可写为

```go

s := "hello"
s = "c" + s[1:] // 字符串虽不能更改但可以做切片操作
fmt.Printf("%s\n", s)

```

如果要声明一个多行的字符串怎么办

```go

m := `hello
    world`   

```

**(``)** 括起的字符串为Raw字符串, 即字符串在代码中的形式就是打印时的形式, 它没有字符转义, 换行也将原样输出

#### 错误类型
Go语言内置有一个error类型, 专门用来处理错误信息, Go语言的package里面还专门有一个包errors来处理错误

```go

err := errors.New("emit macho dwarf: elf header corrputed")
if err != nil {
	fmt.Printf(err)
}
```