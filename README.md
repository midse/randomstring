# randomstring

This package generates a random string from a pattern.

## About this project



It is based on the Perl module String::Random by Steven Pritchard <steve@silug.org> (now maintained by Shlomi Fish).

You can read more about this Perl module here -> http://search.cpan.org/~shlomif/String-Random-0.28/lib/String/Random.pm


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
'''
