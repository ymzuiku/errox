# errox

In debug env, print error file line

errox can print error like:

```bash
/Users/xxx/Documents/work/go-github/errox/cmd/temp.go:16
Root error, /Users/xxx/Documents/work/go-github/errox/cmd/temp.go:10
Sub error
```

## Install

```bash
go get github.com/ymzuiku/errox
```

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
/Users/xxx/Documents/work/go-github/errox/cmd/temp.go:16
Root error, /Users/xxx/Documents/work/go-github/errox/cmd/temp.go:10
Sub error
```

If errox.Debug = false, it print:

```bash
Root error, Sub error
```
