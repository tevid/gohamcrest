tug - Golang测试工具
================

[![Build Status](https://travis-ci.org/zouyx/tug.svg?branch=master)](https://travis-ci.org/zouyx/tug)
[![codebeat badge](https://codebeat.co/badges/cd82fdce-9938-478b-96de-4b3b2e122df7)](https://codebeat.co/projects/github-com-zouyx-tug-master)
[![Coverage Status](https://coveralls.io/repos/github/zouyx/tug/badge.svg?branch=master)](https://coveralls.io/github/zouyx/tug?branch=master)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![GoDoc](http://godoc.org/github.com/zouyx/tug?status.svg)](http://godoc.org/github.com/zouyx/tug)
[![GitHub release](https://img.shields.io/github/release/zouyx/tug.svg)](https://github.com/zouyx/tug/releases)

* tug是一个匹配的库，可在创建测试时灵活的表达意图。当然也能被用于其他用途。

如何安装
------------

### 安装go环境

[请点我](http://golang.org/doc/install.html) 。安装完，就可以享受你的测试之旅。

### 下载依赖包

``` shell
gopm get github.com/zouyx/tug -v -g
```

或者

``` shell
go get -u github.com/zouyx/tug
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

import . "github.com/zouyx/tug"

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

更多使用方式,请查看项目中的test case.更多信息请参考 [Wiki](https://github.com/zouyx/tug/wiki) .

# 如何贡献
  * Source Code: https://github.com/zouyx/tug/
  * Issue Tracker: https://github.com/zouyx/tug/issues
  
# 许可证
The project is licensed under the [Apache 2 license](https://github.com/zouyx/tug/blob/master/LICENSE).
