# Elign

[![Build Status](https://travis-ci.org/visig9/elign.svg?branch=master)](https://travis-ci.org/visig9/elign)

Align the string with east asian character set in golang.



## Usage

```go
package main

import (
    "fmt"
    "gitlab.com/visig/elign"
)

func main() {
	data := []string{
		"世界上",
		"只有 10 種人",
		"懂二進位和不懂二進位的",
	}
	e := elign.Default(0).AdjustWidth(data...)

	for _, d := range data {
		fmt.Printf("|%v|\n", e.Right(d))
	}

	// Output:
	// |                世界上|
	// |          只有 10 種人|
	// |懂二進位和不懂二進位的|
}
```

Please check `godoc` for more information.

## Download

```bash
go get github.com/visig9/elign
```


## License

MIT
