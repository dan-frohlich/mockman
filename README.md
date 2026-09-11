# mockman

A simple Go mock generator. Parses a source file for interfaces (via `go/ast`)
and emits a hand-written-style mock struct for each: func-typed fields plus
methods that call the field if set and `panic` otherwise.

## Usage

```bash
go run . -s <source> -d <dest.go> -p <package> [-f <name filter>]
```

| Flag | Meaning                                              |
|------|----------------------------------------------------|
| `-s` | source file or directory to scan for interfaces     |
| `-d` | destination file to write                           |
| `-p` | package name for the generated file                 |
| `-f` | optional interface-name filter pattern              |

## Example

`examples/foo.go` drives generation with `go:generate` directives:

```go
//go:generate go run .. -s . -p examples -f Foo -d ./foo_mock.go
type Foo interface {
    Bar() string
    Baz(v int) error
    Qux(name string, age int) (bool, error)
}
```

produces `examples/foo_mock.go`:

```go
type FooMock struct {
    bar func() string
    baz func(v int) error
    qux func(name string, age int) (bool, error)
}

func (mock *FooMock) Bar() string {
    if mock.bar != nil {
        return mock.bar()
    }
    panic("Method Bar not implemented in FooMock")
}
// ...
```

Run `go generate ./examples/` to regenerate.
