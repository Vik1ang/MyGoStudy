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
