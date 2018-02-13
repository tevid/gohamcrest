gomatcher - Powerful matcher for Golang
================

[![Build Status](https://travis-ci.org/zouyx/gomatcher.svg?branch=master)](https://travis-ci.org/zouyx/gomatcher)
[![codebeat badge](https://codebeat.co/badges/0b106c76-9761-4c98-9daa-6b123b5f2fa6)](https://codebeat.co/projects/github-com-zouyx-gomatcher-master)
[![Coverage Status](https://coveralls.io/repos/github/zouyx/gomatcher/badge.svg?branch=master)](https://coveralls.io/github/zouyx/gomatcher?branch=master)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![GoDoc](http://godoc.org/github.com/zouyx/gomatcher?status.svg)](http://godoc.org/github.com/zouyx/gomatcher)
[![GitHub release](https://img.shields.io/github/release/zouyx/gomatcher.svg)](https://github.com/zouyx/gomatcher/releases)

**gomatcher** is a powerful matcher for Golang testing designed for easy to write test case. 

Simple, readable and humanized, also easy to extended. 

Help you become an excellent Golang Software Engineer is my believe.

[Chinese doc](README_CN.md)

Installation
------------

Install Go Env [Getting Started](http://golang.org/doc/install.html) .Once finished,then enjoy yourself.

``` shell
gopm get github.com/zouyx/gomatcher -v -g
```

Or

``` shell
go get -u github.com/zouyx/gomatcher
```


*PS*: Prefer Golang 1.6+

# Features
* Assert.List
* Assert.Object
* Assert.String

More feature please respect.

# Usage

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

More usage,check test case please.More info check [Wiki](https://github.com/zouyx/gomatcher/wiki) .

If you think the tool is good or have any problems,must let me know, contact me by email when feel free or create a issue.

# Contribution
  * Source Code: https://github.com/zouyx/gomatcher/
  * Issue Tracker: https://github.com/zouyx/gomatcher/issues
  
# License
The project is licensed under the [Apache 2 license](https://github.com/zouyx/gomatcher/blob/master/LICENSE).
