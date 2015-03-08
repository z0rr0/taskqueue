# Package "taskqueue"

[![GoDoc](https://godoc.org/github.com/z0rr0/taskqueue?status.svg)](https://godoc.org/github.com/z0rr0/taskqueue) [![LGPL License](http://img.shields.io/badge/license-LGPLv3-blue.svg)](https://www.gnu.org/licenses/lgpl-3.0.txt) [![Build Status](https://travis-ci.org/z0rr0/taskqueue.svg?branch=master)](https://travis-ci.org/z0rr0/taskqueue) [![Coverage](http://img.shields.io/badge/coverage-100%-brightgreen.svg)](https://travis-ci.org/z0rr0/taskqueue)


It is a Go package to run/stop periodic tasks. It uses 2 queues to handle running and sleeping tasks. The method [TestStart](https://github.com/z0rr0/taskqueue/blob/master/taskqueue_test.go#L87) from the file [taskqueue_test.go](https://github.com/z0rr0/taskqueue/blob/master/taskqueue_test.go) contains an example of this library usage.

<img src="https://raw.githubusercontent.com/z0rr0/taskqueue/master/img.png" title="image">

### Dependencies

Standard [Go library](http://golang.org/pkg/).

### Design guidelines

There are recommended style guides:

* [The Go Programming Language Specification](https://golang.org/ref/spec)
* [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments).

Check using [go-lint](http://go-lint.appspot.com/github.com/z0rr0/taskqueue) tool.

### Testing

Standard Go testing way:

```shell
cd $GOPATH/src/github.com/z0rr0/taskqueue
go test -v -cover
```

---

*This source code is governed by a [LGPLv3](https://www.gnu.org/licenses/lgpl-3.0.txt) license that can be found in the [LICENSE](https://github.com/z0rr0/taskqueue/blob/master/LICENSE) file.*

<img src="https://www.gnu.org/graphics/lgplv3-147x51.png" title="LGPLv3 logo">
