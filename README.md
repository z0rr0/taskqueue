# Package "taskqueue"

It is an easy library to run/stop a queue of periodic tasks.

### Dependencies

Standard [Go library](http://golang.org/pkg/).

### Design guidelines

There are recommended style guides:

* [The Go Programming Language Specification](https://golang.org/ref/spec)
* [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments).

A compliance with the second style guide can be checked using [go-lint](http://go-lint.appspot.com/github.com/z0rr0/taskqueue) tool.

### Testing

Use standard Go testing mechanism:

```shell
cd $GOPATH/src/github.com/z0rr0/taskqueue
go test
```

Profile coverage is save in [coverage.out](https://github.com/z0rr0/taskqueue/blob/master/coverage.out) file. There is a [nice article](http://blog.golang.org/cover) about tests covering.

### License

This source code is governed by a [LGPLv3](https://www.gnu.org/licenses/lgpl-3.0.txt) license that can be found in the [LICENSE](https://github.com/z0rr0/taskqueue/blob/master/LICENSE) file.

<img src="https://www.gnu.org/graphics/lgplv3-147x51.png" title="LGPLv3 logo">
