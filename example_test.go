package ring_test

import (
	"fmt"

	"github.com/rabingaire/ring"
)

func ExampleRing() {
	r, err := ring.New[string](0)
	fmt.Println(r, err)
	r, err = ring.New[string](5)
	r.Put("A")
	fmt.Println(r.Size())
	fmt.Println(r.Capacity())
	r.Put("B")
	r.Put("C")
	r.Put("D")
	fmt.Println(r.Size())
	fmt.Println(r.Capacity())
	r.Put("E")
	r.Put("F")
	fmt.Println(r.Size())
	fmt.Println(r.Capacity())
	fmt.Println(r.Get())
	fmt.Println(r.Size())
	fmt.Println(r.Capacity())
	fmt.Println(r.Get())
	fmt.Println(r.Get())
	fmt.Println(r.Size())
	fmt.Println(r.Capacity())
	r.Put("G")
	r.Put("H")
	fmt.Println(r.Get())
	fmt.Println(r.Size())
	fmt.Println(r.Capacity())
	fmt.Println(r.Get())
	fmt.Println(r.Get())
	fmt.Println(r.Get())
	fmt.Println(r.Get())
	// Output: <nil> buffer capacity must be greater than zero
	// 1
	// 5
	// 4
	// 5
	// 5
	// 5
	// B <nil>
	// 4
	// 5
	// C <nil>
	// D <nil>
	// 2
	// 5
	// E <nil>
	// 3
	// 5
	// F <nil>
	// G <nil>
	// H <nil>
	//  buffer is empty
}
