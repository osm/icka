# flen

A simple command flag and environment variable parsing library

## Example

```sh
$ cat > test.go <<EOF
package main

import (
        "flag"
        "fmt"

        "github.com/osm/flen"
)

func main() {
        name := flag.String("name", "foo", "A name")
        flen.SetEnvPrefix("TEST")
        flen.Parse()

        fmt.Println("hello", *name)
}
EOF

$ go run test.go
hello foo

$ TEST_NAME=bar go run test.go
hello bar

$ go run test.go -name baz
hello baz
```
