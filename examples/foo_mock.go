package examples

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

func (mock *FooMock) Baz(v int) error {
	if mock.baz != nil {
		return mock.baz(v)
	}
	panic("Method Baz not implemented in FooMock")
}

func (mock *FooMock) Qux(name string, age int) (bool, error) {
	if mock.qux != nil {
		return mock.qux(name, age)
	}
	panic("Method Qux not implemented in FooMock")
}
