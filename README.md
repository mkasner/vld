#vld

Validator engine for Golang. Validate struct fields or write custom validations as functions.
Chain them together and make your code cleaner and easier to reason.2

#Installation

Make sure that Go is installed on your computer. Type the following command in your terminal:
```
go get github.com/mkasner/vld
```
#Example
Example struct we want to validate with the following rules:
- Field Kind is required and must not be longer then 10 characters
- Feld year is required and must not be less than 1900 and greater than 3000

```go
type Wine struct {
  Kind string
  Year int
}
```
The ususal way would be something like this
```go

func (w *Wine) Validate() error {
    if w.Kind == "" {
      return errors.New("Wine.Kind required")
    }
    if len(w.Kind) > 10 {
      return errors.New("Wine.Kind must be smaller than 10 characters")
    }
    if w.Year == 0 {
      return errors.New("Wine.Year is required")
    }
    if w.Year < 1900 {
      return errors.New("Wine.Year must be greater than 1900")
    }
    if w.Year > 3000 {
      return errors.New("Wine.Year must be smaller than 3000")
    }
    return nil
}
func isValid(w *Wine) bool {
      if err := w.Validate(); err != nil {
        return true
      }
      return false
}
```
Or with the help of vld utility
```go
//go:generate vld -type=Wine
package cellar

type Wine struct {
  Kind string `vld:"req,maxlen=10"`
  Year int  `vld:"req,maxint=3000,minint=1900"`
}

func isValid(w Wine) bool {
    if err := vld.New().Add(w).Validate(); err != nil {
      return true
    }
    return false
}
```
It generates struct validator for you, so you can just call validate method on struct and get results.
#Thanks
- [github.com/asaskevich/govalidator](https://github.com/asaskevich/govalidator)  - lots of documented and tested validations
- [https://godoc.org/golang.org/x/tools/cmd/stringer](https://godoc.org/golang.org/x/tools/cmd/stringer) - intro into AST package and code generation