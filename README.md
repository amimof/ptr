# ptr
Library for dealing with pointers in Go


## Install
`
go get github.com/amimof/ptr
`

## Use

```Go
import (
  "fmt"
  "github.com/amimof/ptr"
)

func main() {
  s := ptr.StrPtr("This is a string")
  fmt.Printf("%v", s)
}
```