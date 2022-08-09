package simple

type Foo struct {
}

// provider/constructor
func NewFoo() *Foo {
	return &Foo{}
}

type Bar struct {
}

// func provider
func NewBar() *Bar {
	return &Bar{}
}

type FooBar struct {
	*Foo
	*Bar
}

// sblmnya kita langsung buat provider dlm bntk func, now dlm bntk struct di injector.go
