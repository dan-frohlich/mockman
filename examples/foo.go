//go:generate echo "starting generation"
//go:generate go run .. -s . -p examples -f Foo -d ./foo_mock.go
//go:generate gofmt -w ./foo_mock.go
//go:generate echo "generation complete"
package examples

type Foo interface {
	Bar() string
	Baz(v int) error
	Qux(name string, age int) (bool, error)
}
