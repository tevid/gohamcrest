gohamcrest - Hamcrest-style matchers for Golang testing
================

[![Build Status](https://travis-ci.org/zouyx/gohamcrest.svg?branch=master)](https://travis-ci.org/zouyx/gohamcrest)
[![codebeat badge](https://codebeat.co/badges/6b5ab21f-16a7-457c-b247-ba7d13bda3eb)](https://codebeat.co/projects/github-com-zouyx-gohamcrest-master)
[![Coverage Status](https://coveralls.io/repos/github/zouyx/gohamcrest/badge.svg?branch=master)](https://coveralls.io/github/zouyx/gohamcrest?branch=master)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![GoDoc](http://godoc.org/github.com/zouyx/gohamcrest?status.svg)](http://godoc.org/github.com/zouyx/gohamcrest)
[![GitHub release](https://img.shields.io/github/release/zouyx/gohamcrest.svg)](https://github.com/zouyx/gohamcrest/releases)

**gohamcrest** is Hamcrest-style matchers for Golang testing designed for easy to write test case.

Simple, readable and humanized, also easy to extended. 

Help you become an excellent Golang Software Engineer is my believe.

[Chinese doc](README_CN.md)

Installation
------------

Install Go Env [Getting Started](http://golang.org/doc/install.html) .Once finished,then enjoy yourself.

``` shell
gopm get github.com/zouyx/gohamcrest -v -g
```

Or

``` shell
go get -u github.com/zouyx/gohamcrest
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

More usage,check test case please.More info check [Wiki](https://github.com/zouyx/gohamcrest/wiki) .

If you think the tool is good or have any problems,must let me know, contact me by email when feel free or create a issue.

# Contribution
  * Source Code: https://github.com/zouyx/gohamcrest/
  * Issue Tracker: https://github.com/zouyx/gohamcrest/issues
  
# License
The project is licensed under the [Apache 2 license](https://github.com/zouyx/gohamcrest/blob/master/LICENSE).
