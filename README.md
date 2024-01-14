# gcode
a generate library for go code 
# demo
```go
	code := `package aa`
	bs, err := AddImport([]byte(code), "arr", "github.com/lingdor/magicarray")
	if err != nil {
		panic(err)
	}
    fmt.Println(string(bs))
```
output:
```go
package aa
import(
arr "github.com/lingdor/magicarray"
)
```