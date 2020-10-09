# constructor [![PkgGoDev](https://pkg.go.dev/badge/github.com/gostaticanalysis/constructor)](https://pkg.go.dev/github.com/gostaticanalysis/constructor)

`constructor` constructor reports whether name of a constructor like function does not begin "New".

A constructor like function:

* Requests named struct type or its pointer
* Requests one or two value, last value must be an error
* Exported function and it is not method

```go
type T1 struct{}

func NewT1() (t *T1) { return } // OK

func CreateT1() (t *T1) { return } // want `name of a constructor like function must begin "New"`
```
