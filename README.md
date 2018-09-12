# JSON parser for Golang

Simple JSON parser for golang.
This is for private study of goyacc.

## Example

```console
$ go get github.com/uphy/go-json
```

```golang
import "github.com/uphy/go-json"

func main(){
    v, _ := ParseObject(`{"a":1}`)
    fmt.Println(v)
}
```