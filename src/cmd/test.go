package main

import "fmt"

//encapsulation && inheritance
/*type Foo struct {
	baz string
}

type Bar struct {
	Foo
}

func (f *Foo) echo() {
	fmt.Println(f.baz)
}

func main() {
	f := Foo{"hello encapsulation"}
	f.echo()

	b := Bar{Foo{"hello inheritance"}}
	b.echo()
}
*/

//polymorphism

type Foo interface {
	 qux()
}

type Bar struct {}
type Baz struct {}

func (b *Bar) qux() {
	fmt.Println("bar polymorphism")
}

func (b *Baz) qux() {
	fmt.Println("baz polymorphism")
}

func main()  {
	var f Foo
	f = &Bar{}
	f.qux()
	fmt.Println(&f)
	fmt.Println(*&f)
	fmt.Println(f)

	f = new(Baz)
	f.qux()
	fmt.Println(*new(Baz))
	fmt.Println(new(Baz))
}