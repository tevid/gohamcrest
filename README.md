tug - Test utils for Golang
================

[![Build Status](https://travis-ci.org/zouyx/tug.svg?branch=master)](https://travis-ci.org/zouyx/tug)
[![codebeat badge](https://codebeat.co/badges/cd82fdce-9938-478b-96de-4b3b2e122df7)](https://codebeat.co/projects/github-com-zouyx-tug-master)
[![Coverage Status](https://coveralls.io/repos/github/zouyx/tug/badge.svg?branch=master)](https://coveralls.io/github/zouyx/tug?branch=master)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![GoDoc](http://godoc.org/github.com/zouyx/tug?status.svg)](http://godoc.org/github.com/zouyx/tug)
[![GitHub release](https://img.shields.io/github/release/zouyx/tug.svg)](https://github.com/zouyx/tug/releases)

Test utils for Golang，use in go test only.[Chinese doc](README_CN.md)

Installation
------------

Install Go Env [Getting Started](http://golang.org/doc/install.html) .Once finished,then enjoy yourself.

``` shell
gopm get github.com/zouyx/tug -v -g
```

Or

``` shell
go get -u github.com/zouyx/tug
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

More usage,check test case please.More info check [Wiki](https://github.com/zouyx/tug/wiki) .

If you think the tool is good or have any problems,must let me know, contact me by email when feel free or create a issue.

# Contribution
  * Source Code: https://github.com/zouyx/tug/
  * Issue Tracker: https://github.com/zouyx/tug/issues
  
# License
The project is licensed under the [Apache 2 license](https://github.com/zouyx/tug/blob/master/LICENSE).
