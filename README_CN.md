gohamcrest - Golang的断言库
================

[![Build Status](https://travis-ci.org/zouyx/gohamcrest.svg?branch=master)](https://travis-ci.org/zouyx/gohamcrest)
[![codebeat badge](https://codebeat.co/badges/6b5ab21f-16a7-457c-b247-ba7d13bda3eb)](https://codebeat.co/projects/github-com-zouyx-gohamcrest-master)
[![Coverage Status](https://coveralls.io/repos/github/zouyx/gohamcrest/badge.svg?branch=master)](https://coveralls.io/github/zouyx/gohamcrest?branch=master)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![GoDoc](http://godoc.org/github.com/zouyx/gohamcrest?status.svg)](http://godoc.org/github.com/zouyx/gohamcrest)
[![GitHub release](https://img.shields.io/github/release/zouyx/gohamcrest.svg)](https://github.com/zouyx/gohamcrest/releases)

**gohamcrest**是一个断言库，该库用于更方便的写测试用例。

可在创建测试时灵活的表达意图，当然也能被用于其他用途。

助君成为一个出色的Golang软件工程师是我的景愿。

如何安装
------------

### 安装go环境

[请点我](http://golang.org/doc/install.html) 。安装完，就可以享受你的测试之旅。

### 下载依赖包

``` shell
gopm get github.com/zouyx/gohamcrest -v -g
```

或者

``` shell
go get -u github.com/zouyx/gohamcrest
```


*PS*: 最好使用 Golang 1.6+

# 功能
* 校验List
* 校验Object
* 校验数字，如int，float
* 校验String

更多功能，敬请期待.

# 如何使用

### 导入包

import . "github.com/zouyx/gohamcrest"

### 使用

- Assert equal

``` go
func TestEqual(t *testing.T) {
	Assert(t,2,Equal(2))
	Assert(t,"joe",Equal("joe"))
}
```

- Assert not equal

``` go
func TestNotEqual(t *testing.T) {
	Assert(t,2,NotEqual(3))
	Assert(t,"joe",NotEqual("joe1"))
}
```

or 

``` go
func TestNotEqual(t *testing.T) {
	Assert(t,2,Not(Equal(3)))
	Assert(t,"joe",Not(Equal("joe1")))
}
```

更多使用方式,请查看项目中的test case.更多信息请参考 [Wiki](https://github.com/zouyx/gohamcrest/wiki) .

# 如何贡献
  * Source Code: https://github.com/zouyx/gohamcrest/
  * Issue Tracker: https://github.com/zouyx/gohamcrest/issues
  
# 许可证
The project is licensed under the [Apache 2 license](https://github.com/zouyx/gohamcrest/blob/master/LICENSE).
