# randomstring

[![Build Status](https://drone.io/github.com/midse/randomstring/status.png)](https://drone.io/github.com/midse/randomstring/latest) [![Coverage Status](https://coveralls.io/repos/midse/randomstring/badge.svg?branch=master&service=github)](https://coveralls.io/github/midse/randomstring?branch=master)

This package generates a random string from a pattern.

## Install this package

```
$ go get github.com/midse/randomstring
```

## Basic usage

```go
import (
  "fmt"
  "github.com/midse/randomstring"
)

func main() {
    result := randomstring.FromRegex("\\#\\d\\d\\d")
    fmt.Println(result)
}
```


## About this project

Code needs some cleanup and improvements. It was created to implement a basic fuzzer.
By the way, this module deals only with a subset of regular expressions.

Feel free to contribute.

It's mainly based on the Perl module String::Random by Steven Pritchard <steve@silug.org> (now maintained by Shlomi Fish).

You can read more about this Perl module here -> http://search.cpan.org/~shlomif/String-Random-0.28/lib/String/Random.pm
