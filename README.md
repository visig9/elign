# Elign

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
go get gitlab.com/visig/elign
```


## License

MIT
