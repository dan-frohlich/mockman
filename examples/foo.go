package examples

type Foo interface {
	Bar() string
	Baz(int) error
	Qux(name string, age int) (bool, error)
}
