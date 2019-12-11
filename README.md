# 说明

## 可视化网站

### 地址

[https://502408764.github.io/#/](https://502408764.github.io/#/)

### 项目地址

[https://github.com/543877815/os-visualization](https://github.com/543877815/os-visualization)

## 序列检查程序

如果您想查看您的程序的输出结果与正确结果不同的地方，可以使用`result_check_windows.exe`和`result_check_linux`序列检查程序，它可以将您的输出与答案对比，高亮出不同的地方：

* 序列检查程序会帮助您自动运行运行所有示例
* 序列检查程序目前只支持运行`可执行程序`和`.py`文件

**在Windows环境下，推荐使用Windows Terminal或其他现代化的终端，否则标记颜色不会被正常显示*

### 如果您的实验文件是可执行文件

使用方式如下：

①将您已经编译好的实验文件放在本仓库的根目录下

②打开终端，`cd`进入该仓库文件夹，注意第一个参数是您的可执行文件名（请直接传递可执行文件名，不带`./`）

* Windows

  ```powershell
  D:>\os_test_shell\result_check_windows.exe your_file_name.exe
  ```

* Linux

  ```powershell
  linux@someone-linux~$./os-test-shell/result_check_linux your_file_name
  ```

### 如果您的实验文件为.py文件

使用方式如下：

①将您的实验文件放在本仓库的根目录下

②打开终端，`cd`进入该仓库文件夹，注意第一个参数是您的`.py`文件名（请直接传递可执行文件名，不带`./`），第二个参数根据不同的系统，需要输入不同的参数

* Windows

  ```powershell
  D:>\os_test_shell\result_check_windows.exe xxx.py python
  ```

* Linux

  ```powershell
  # 如果您使用python3编写，请执行下面这条命令
  linux@someone-pc~$./os-test-shell/result_check_linux xxx.py python3
  # 如果您使用python2编写，请执行下面这条命令
  linux@someone-pc~$./os-test-shell/result_check_linux xxx.py python2
  ```

*源代码在根目录的source_code文件夹

* 如果您的程序运行通过一个测试，终端上会显示：

  ![correct](https://raw.githubusercontent.com/joexu01/joexu01.github.io/master/result_correct.png)

* 如果您的程序运行没有通过测试，终端上会将错误的输出标记：

  ![incorrect](https://raw.githubusercontent.com/joexu01/joexu01.github.io/master/incorrect_answer.png)

* 如果您的程序通过了全部测试，终端上会显示：

  ![full](https://raw.githubusercontent.com/joexu01/joexu01.github.io/master/full.png)

* 请注意，序号下方第一行是您的程序的输出，第二行是预期的正确输出，所有超过正确输出长度的输出都会被忽略掉

## 提问issue的正确方式

1. 给出输入序列
2. 给出期望输出和实际输出序列
3. 指明哪一行输出有问题
4. 简单地分析一下？
5. 如果是可视化网站的问题或者要提PR，请@[李逢君](https://github.com/543877815)
6. 如果是序列检查程序的问题，请@[许思博](https://github.com/joexu01)
7. 欢迎fork

## 实验文件的参数接收

推荐您将测试用例作命令行运行的参数传递，如，您的实验文件为main.exe，您应该将测试用例的文件名作为第一个参数输入命令行

```powershell
./main.exe test_shell.txt
```

## 输出建议

不要换行！

不要换行！

不要换行！

严格按照实验指导书的样式

## 具体用例

### 0.txt

#### 输入

```
cr x 1
cr p 1
cr q 1
cr r 1
to
req R2 1
to
req R3 3
to
req R4 3
to
to
req R3 1
req R4 2
req R2 2
to
de q
to
to
```

#### 输出

```
init x x x x p p q q r r x p q r x x x p x 
```

### 1.txt

#### 输入

```
cr A 1
cr B 1
cr C 1
to
cr D 1
cr E 1
to
cr F 1
req R1 1
req R2 2
to
req R2 1
req R3 3
to
req R4 4
to
req R3 2
to
rel R2 1
to
rel R3 2
to
to
req R3 3
de B
to
to
to
```

#### 输出

```
init A A A B B B C C C C A D D E E B F C C D D E F A A C F A 
```

### 2.txt

#### 输入

```
cr A 1
cr B 1
cr C 1
req R1 1
to
cr D 1
req R2 2
cr E 2
req R2 1
to
to
to
rel R2 1
de B
to
to
```

#### 输出

```
init A A A A B B B E C A D B E C A C 
```

### 3.txt

#### 输入

```
cr A 1
cr B 1
cr C 1
to
cr D 1
cr E 1
cr F 1
to
to
to
req R1 1
req R2 1
to
req R2 1
to
req R3 3
req R4 3
req R4 3
to
req R1 1
cr G 2
req R1 1
de B
req R3 2
cr H 2
cr I 2
to
req R3 3
req R3 2
to
rel R3 1
to
```

#### 输出

```
init A A A B B B B C A D D D E E F F F B C A G D A A H H I H C A A C
```

### 4.txt

#### 输入

```
cr x 1
cr p 1
cr q 1
cr r 1
to
to
to
req R1 1
to
req R2 1
to
req R3 2
to
to
rel R1 1
to
req R3 3
de p
to
```

#### 输出

```
init x x x x p q r r x x p p q r r x p q r
```

### 5.txt

#### 输入

```
cr a 1
cr b 1
cr c 1
cr d 1
to
cr f 1
req R1 1
to
to
to
cr e 2
req R1 1
to
de b
req R1 1
to
to
to
to
to
```

#### 输出

```
init a a a a b b b c d a e f b e c d a c d a
```

### 6.txt

#### 输入

```
cr a 1
cr b 1
to
cr c 1
cr d 1
to
req R1 1
to
to
req R2 2
to
de b
req R3 1
to
```

输出

```
init a a b b b a a c d d b a a a
```

### 7.txt

#### 输入

```
cr z 1
cr x 1
cr c 1
to
req R3 3
cr v 1
to
to
req R3 1
to
req R1 1
to
req R1 1
de x
to
```

#### 输出

```
init z z z x x x c z v x x c v z c
```

### 8.txt

#### 输入

```
cr a 1
cr s 1
to
cr d 1
req R2 2
cr f 1
to
to
req R2 1
to
de s
to
req R2 1
```

#### 输出

```
init a a s s s s a d f s a a a
```

### 9.txt

#### 输入

```
cr x 1
cr y 1
req R2 2
to
cr z 1
cr m 1
req R1 1
to
req R2 2
de x
to
```

#### 输出

```
init x x x y y y y x z init init
```

## 贡献

- 0 实验指导书
- 1、2 、5[万志文](https://github.com/JXhacker)
- 3 [李逢君](https://github.com/543877815)
- 4 黄晔熙
- 6~9 [李若欣](https://github.com/Karenlrx)
- shell [向尉](https://github.com/SwordAndTea)
- 序列检查 [许思博](https://github.com/joexu01)
