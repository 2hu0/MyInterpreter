# MyInterpreter

### 前言

众所周知，编译原理是程序员的三大浪漫之一。本学期我们也上了这门课，但上课时诸多概念让我眼花缭乱，也不知道这些概念究竟对应着什么。于是我准备用实战的方式来进行课堂之外的学习，主要参考资料有：

* 《Writing An Interpreter In Go》

* 极客时间专栏《编译原理之美》

* 极客时间专栏《编译原理实战》

正好最近学习了Go的语法。也能巩固一下。Go语言十分简洁，包括Go语言提供的测试框架也很好用，比之前用C++写课程实验要友好一些。

#### 解释器

首先我先尝试实现一个解释器，它将通过解析源代码，构建抽象语法树`(AST)`，然后对这棵树进行求值。这类解释器也被称为树遍历`(tree-walking)`解释器，因为工作方式是在AST上遍历并且解释。

本解释器内容包括

* 词法分析器
* 语法分析器
* 抽象语法树
* 求值器

### Monkey编程语言

每个解释器都用来解析一种特定的编程语言，本项目中我构建的是**Monkey**语言的解释器。它是存在于书籍上的语言，为什么叫它Monkey呢？作者说是因为猴子是一个聪明有趣的动物，正如同编译器一样。

语言特性如下

* 类C语言语法
* 变量类型绑定
* 整型和bool类型
* 算术运算
* 内置函数
* 头等函数和高阶函数
* 闭包
* string结构
* 哈希数据结构
* 数组数据结构



**使用样例如下**

```markdown
let name = "2hu0";
let age = 18;
let res = (10 + 2 ) / 3;
let add = fn(a,b){ return a + b };
let add = fn(a,b){ a + b };
add(1,2);
let arr = {1,2,3};
let info = {"name":"2hu0","age":18};
let twice = fn(f,x) {
 return f(f(x));
}
```


