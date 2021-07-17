# errox

In debug env, print error file line

errox can print error like:

```bash
/Users/xxx/Documents/work/go-github/errox/cmd/temp.go:16 (Method main.errorB)
|- Root error, /Users/xxx/Documents/work/go-github/errox/cmd/temp.go:10 (Method main.errorA)
|- Sub error
```

## Install

```bash
go get github.com/ymzuiku/errox
```

## APIs

- `errox.Debug` : Change debug bool
- `errox.New(string) error` : Create a error
- `errox.Wrap(error) error` : Add stack info to error when errox.Debug = true
- `errox.Errorf(string, args ...interface{}) error` : Add stack info and format error when errox.Debug = true
- `errox.Line() string` : Make stack info string when errox.Debug = true

## Use

```go
package main

import (
	"fmt"

	"github.com/ymzuiku/errox"
)

func errorA() error {
	err := errox.New("Sub error")
	return err
}

func errorB() error {
	errA := errorA()
	err := errox.Errorf("Root error, %w", errA)

	return err
}

func main() {
	// Dont't set production debug = true
	errox.Debug = true
	err := errorB()
	fmt.Println(err)
}
```

If errox.Debug = true, it print:

```bash
/Users/xxx/Documents/work/go-github/errox/cmd/temp.go:16 (Method main.errorB)
|- Root error, /Users/xxx/Documents/work/go-github/errox/cmd/temp.go:10 (Method main.errorA)
|- Sub error
```

If errox.Debug = false, it print:

```bash
Root error, Sub error
```
